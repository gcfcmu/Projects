
import pandas as pd
import seaborn
import numpy as np
import scipy.stats.mstats as stats
import matplotlib.pyplot as plt
import sklearn.decomposition as skd


def ParseFromDict(data, dictionary):

    groupToDf = dict()
    parsed = data.copy()

    for groupName in dictionary.keys():

        parsed = data.copy()

        for conditions in range(len(dictionary[groupName])):

            parsed = parsed[parsed[dictionary[groupName][conditions][0]] == dictionary[groupName][conditions][1]]

        groupToDf[groupName] = parsed

    return groupToDf

def MeltAndSplit(data, geneColName = "TargetName", names = ["Slide", "AOI", "Segment"]):

    geneColName = VerifyGeneColumn(data, geneColName)

    spdf = pd.melt(data, id_vars=geneColName, value_vars=data.keys())
    spdf = spdf.rename(columns = \
                   {geneColName: "Gene", "variable": "ROI", "value" : "Counts"})

    splitted = [spdf["ROI"].values[i].split("|") for i in range(len(spdf["ROI"]))]
    splitted = np.array(splitted)


    for i in range(len(names)):
        spdf[names[i]] = splitted[:, i]

    return spdf


def SignalToNoise(data, negativeProbeName = "Negative Probe", geneColName = "TargetName", return_log = True):

    geneColName = VerifyGeneColumn(data, geneColName)

    out = data.copy()

    noiseIndex = np.where(out[geneColName] == "Negative Probe")[0][0]

    for i in out.keys():

        if not (isinstance(out[i][0], str)):
            if return_log:
                out[i] = np.log1p(out[i] / out[i][noiseIndex])
            else:
                out[i] = out[i] / out[i][noiseIndex]

    return(out)






# q3 performs quantile normalization on the wide-form (matrix-like) data. The argument copy can be used to create a copy
# of the data such that the function does not alter the original data. With that said, the default argument is false because
# there is little use to keep a copy of the original data. Note also that it's necessary to know what the column containing
# the gene names is called, as well as what the negative probe (background/noise) is called in the gene column. The function
# must not use the negative probe to calculate the third quantile for normalization, as it is not technically a gene count.
# Because of this, a copy of the data is made to ensure that this value is not used for the math, but is retained in the output.
def q3(data, geneColName = "TargetName", negativeProbeName = "Negative Probe", copy = False):

    geneColName = VerifyGeneColumn(data, geneColName)

    if copy:
        return(q3Copy(data, geneColName = "TargetName", negativeProbeName = "Negative Probe"))

    trueData = data.copy()

    genes = data[geneColName]
    noiseIndex = np.where(genes == negativeProbeName)[0][0]

    trueData = trueData.drop(noiseIndex, axis=0)

    quantiles = [0 for i in range(len(trueData.loc[0])-1)]

    j = 0
    for i in trueData.keys():
        if i != geneColName:
            quantiles[j] = np.quantile(trueData[i], 0.75)
            j = j + 1

    geomean = stats.gmean(quantiles)

    j = 0
    for i in data.keys():
        if i != geneColName:
            data[i] = data[i] / quantiles[j] * geomean
            j = j + 1

    return data


def q3Copy(data, geneColName = "TargetName", negativeProbeName = "Negative Probe"):

    copyData = data.copy()
    trueData = data.copy()

    genes = data[geneColName]
    noiseIndex = np.where(genes == negativeProbeName)[0][0]

    trueData = trueData.drop(noiseIndex, axis=0)

    quantiles = [0 for i in range(len(trueData.loc[0])-1)]

    j = 0
    for i in trueData.keys():
        if i != geneColName:
            quantiles[j] = np.quantile(trueData[i], 0.75)
            j = j + 1

    geomean = stats.gmean(quantiles)

    j = 0
    for i in copyData.keys():
        if i != geneColName:
            copyData[i] = copyData[i] / quantiles[j] * geomean
            j = j + 1

    return copyData


# FilterGenes removes genes that are below the negative probe count. If, for a single sample,
# the count is less than the negative probe for some percentage of samples, that gene will be removed.

# Example:
# test = pd.DataFrame({"Gene":["GeneA", "GeneB", "NP"], "Sample1": [2.0, 0.5, 1.0], "Sample2": [3.0, 2.5, 0.9]})
# out = FilterGenes(data = test, percentSamples = 75, geneColName = "Gene", negativeProbeName = "NP")
# print(out)

# Before filtering:
#    Gene  Sample1  Sample2
#0  GeneA      2.0      3.0
#1  GeneB      0.5      2.5
#2     NP      1.0      0.9

# After filtering:
#    Gene  Sample1  Sample2
#0  GeneA      2.0      3.0
#2     NP      1.0      0.9

# In this example, since GeneB is less than the negative probe for sample1, which constitutes 50% of samples, it's not present
# in 75% of samples. So, it is filtered out.

def FilterGenes(data, percentSamples, geneColName = "TargetName", negativeProbeName = "Negative Probe", sdCutoff = 0):

    geneColName = VerifyGeneColumn(data, geneColName)

    # the number of samples times the percent gives the flat number needed
    numSamples = len(data.keys()) - 1
    threshold = numSamples * percentSamples / 100

    # extract the negative probe count from the data
    noiseIndex = np.where(data[geneColName] == negativeProbeName)[0][0]
    noiseAcrossSamples = data.loc[noiseIndex]
    stdevs = np.std(data[data[geneColName] != negativeProbeName]).values
    stdevAbove = (stdevs * sdCutoff) + noiseAcrossSamples[1:]


    # subtract the negative probe count from the actual counts
    # then, filter the dataset such that we only keep genes that have more signal
    # than noise. Note this will filter out the negative probe count too
    minusNoise = data.to_numpy(copy = True)
    minusNoise = minusNoise[:,1:] - stdevAbove.values
    toKeep = [(sum(minusNoise[i] > 0) > threshold) for i in range(len(minusNoise))]
    toKeep[noiseIndex] = True
    out = data[toKeep]



    return(out)

def FilterGenesAlt(data, percentSamples):

    # the number of samples times the percent gives the flat number needed
    numSamples = len(data.keys()) - 1
    threshold = numSamples * percentSamples / 100

    # extract the negative probe count from the data
    noiseIndex = np.where(data["TargetName"] == "Negative Probe")[0][0]
    noiseAcrossSamples = data.loc[noiseIndex]

    # subtract the negative probe count from the actual counts
    # then, filter the dataset such that we only keep genes that have more signal
    # than noise.
    minusNoise = data.to_numpy(copy = True)
    minusNoise = minusNoise[:,1:] - noiseAcrossSamples.values[1:]
    toKeep = [(sum(minusNoise[i] > 0) > threshold) for i in range(len(minusNoise))]
    toKeep[noiseIndex] = True
    out = data[toKeep]

    return(out)

# DropSurfaceArea takes in as input two data frames. data is the matrix-type data containing gene counts for each sample.
# reference is the metadata excel sheet that contains information about each sample. For example, one may filter samples that
# have less than desirable surface area. However, the metric may be changed to filter out samples based on any other numeric
# variable. In addition, the lessThan variable may be used to denote if filtering should be done on samples below or above the
# threshold (i.e. if there is an upper limit). It is also useful to know what the segment names are in the reference dataframe,
# in order to ensure the correct samples are identified.
def DropSurfaceArea(reference, data, threshold, metric = 'AOISurfaceArea', lessThan = True, segmentName = "SegmentDisplayName"):

    sampleNames = []
    index = 0

    if lessThan:
        for i in reference[metric] < threshold:
            if i:
                sampleNames.append(reference[segmentName][index])
            index = index + 1


    else:
        for i in reference[metric] > threshold:
            if i:
                sampleNames.append(reference[segmentName][index])
            index = index + 1


    data = data.drop(sampleNames, axis = 1)

    return(data)



# LogData is a simple function that performs a log1p operation on each entry in the matrix-like dataframe. Copy is used
# if one does not wish to alter the input argument. Note that this function will alter any numeric variable in the data.
# The logarithm is the natural logarithm.

# Example:

# test = pd.DataFrame({"Gene":["GeneA", "GeneB", "GeneC"], "Sample1": [2.0, 0.5, 1.0], "Sample2": [3.0, 2.5, 0.9]})

# Output:

#     Gene   Sample1   Sample2
#0  GeneA  1.098612  1.386294
#1  GeneB  0.405465  1.252763
#2  GeneC  0.693147  0.641854

def LogData(data, copy = False):

    if copy:
        logData = data.copy()
        for i in logData.keys():
            if not (isinstance(logData[i][0], str)):
                logData[i] = np.log1p(logData[i])
        return logData
    else:
        for i in data.keys():
            if not (isinstance(data[i][0], str)):
                data[i] = np.log1p(data[i])
        return data



# LongPanelData manually turns a matrix data into a long-form data. The reason for this is to incorporate the panel dict and
# segment information where appropriate. Data should be in the form of signal to noise, wide format. segmentProperties is
# important as it contains metadata (e.g. segment labels). segLabelName is the column in the metadata that corresponds to
# the segment input. geneColName is again important to specify if not already called "TargetName" in data.

def LongPanelData(data, segmentProperties, targetCountMatrix, panelDict, segment, \
                  geneColName = "TargetName", segLabelName = "SegmentLabel"):

    geneColName = VerifyGeneColumn(data, geneColName)

    panelNames = panelDict.keys()
    booleanPanelIndexes = {pName: \
                       [targetCountMatrix[geneColName].values[i] in panelDict[pName] \
                        for i in range(len(data))] \
                       for pName in panelNames \
                      }

    segmentDataDict = {}

    for segName in segment:
        sList = (segmentProperties[segLabelName] == segName).values
        sList = list(sList)
        sList = [False] + sList
        segmentDataDict[segName] = (targetCountMatrix.loc[:, sList])

    panelData = {}

    datf = 0
    row = 0

    for name in panelNames:
        for seg in segment:
            panelData[name + "|" + seg] = segmentDataDict[seg][booleanPanelIndexes[name]]

    newdf = pd.DataFrame({"Segment": [], "Panel": [], "SampleName": [], \
                          "Replicate": [], "SNR":[]})



    for keys in panelData.keys():
        dataframe = panelData[keys]

        currentPanel = keys.split("|")[0]
        currentSegment = keys.split("|")[1]

        for i in dataframe.keys():
            rep = 0
            for j in dataframe[i]:
                newdf.loc[row] = [currentSegment, currentPanel, i, rep, j]
                rep += 1
                row += 1
        datf += 1

    # Extract the SampleName from the "sample | roi | segment"  column

    alls = []
    for i in newdf["SampleName"]:
        new = i.split(" ")
        alls.append(new[0])

    newdf["Sample"] = alls

    return(newdf)



def PCAData(data, names = ["Slide", "ROI", "Segment"]):

    logTransformed = data.copy()

    LogData(logTransformed)

    pca = skd.PCA(n_components = 2)

    dimReduc = pca.fit_transform(logTransformed.T[1:])

    pcaDf = pd.DataFrame(dimReduc)


    pcaDf = pcaDf.set_axis(["PCA1", "PCA2"], axis = 1)


    splitted = [logTransformed.keys().values[i].split("|") for i in range(1, len(data.keys()))]
    splitted = np.array(splitted)

    for i in range(len(names)):
        pcaDf[names[i]] = splitted[:, i]


    return pcaDf


def PCALabelSamples(pcaDf, sampleLabels):

    progressor = ["" for i in range(len(pcaDf["Slide"]))]

    for label in sampleLabels.keys():
        for substring in sampleLabels[label]:

            for row in range(len(pcaDf["Slide"])):
                if substring in pcaDf["Slide"][row]:
                    progressor[row] = label

    pcaDf["Progressor"] = progressor

    return pcaDf

def FilterLogExpression(data, minExpCutoff = 4, minNumAOI = 3):

    n = len(data[data.keys()[1]])
    outputData = data.copy()
    toKeep = [False for i in range(n)]

    #for all genes, if the gene has expression less than minExpCutoff in more than n - minNumAOI, remove it
    for i in range(n):
        numAcceptable = sum(data.iloc[i][1:] > minExpCutoff)
        if numAcceptable > minNumAOI:
            toKeep[i] = True

    return outputData[toKeep]

# Check if the given string is a column the data. It should contain the column with gene names. If it's not in the data, look
# for the first column that contains a string, and return that column. The input should always be a wide form dataset. This
# function won't work as intended if the data has multiple string columns and the column with the gene names isn't the first
# (e.g. this could happen if the data is incorrectly read as a string). This is a failsafe for when the given gene column name
# is wrong (e.g. omitted, when not called "TargetName").
def VerifyGeneColumn(data, givenString):

    try:
        data[givenString]
    except:
        for row in range(len(data)):
            for i in data.keys():
                if isinstance(data[i][row], str):
                    return(i)
    return(givenString)
