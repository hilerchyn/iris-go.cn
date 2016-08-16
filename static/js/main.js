/**
 * Created by HilerChyn <hilerchyn@gmail.com> on 8/17/16.
 */

$(document).ready(function(){
    //alert(window.location.pathname);

    var pathArray = window.location.pathname.split("/");

    var activeID = pathArray[1];
    if(activeID == ""){
        activeID = "home";
    }

    $("#"+activeID).addClass("active");

    $("img").css("width", "100%");

})