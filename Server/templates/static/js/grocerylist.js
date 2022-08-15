//document.forms['grocerylist']['textField'].readOnly = true;

$(document).ready(function(){
    $("#toggle_readonly").click(function(){
        $("[id=textField]").each(function(){
            if($(this).prop("readonly") == true){
                $(this).prop('readonly', false);
            }else{
                $(this).prop('readonly', true);
            }
        });
    });
});