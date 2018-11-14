function getSinaShortUrl() {
    var str = $("#sinaurl").val();
    console.log(str);
    $.ajax({
        cache : true,
        type : "GET",
        url : "/getSinaShortUrl",
        data : {"url":str},
        async : false,
        error : function() {
            alert("系统异常");
        },
        success : function(data) {
            $('#sUrlDiv').removeClass("hidden")
            if (data.code == 0) {
                var dt = data.data[0].url_short
                $('#sUrl').html(dt);
            }
        }
    });

}