$(function(){
    $('#add-row').click(function(){
       $(".row:first-child").clone().prependTo(".container-fluid");
    });
});
