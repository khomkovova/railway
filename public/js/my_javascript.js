// const Http = new XMLHttpRequest()
// const url='http://127.0.0.1:12345/signin';
// Http.open("POST", url, true);
// // Http.header('"Access-Control-Allow-Origin":"*"');
// // Http.setData( '{"username":"admin", "password":"pass"}');
// Http.send('{"username":"admin", "password":"pass"}');
//
//
// Http.onreadystatechange=(e)=>{
//     console.log(Http.getAllResponseHeaders());
//     console.log("ok");
//     const Http2 = new XMLHttpRequest();
//     const url2='http://127.0.0.1:12345/welcome';
//     Http2.open("GET", url2, true);
//     Http2.xhr().withCredentials(true);
// // Http.header('"Access-Control-Allow-Origin":"*"');
// // Http.setData( '{"username":"admin", "password":"pass"}');
// // Http2.send('{"username":"admin", "password":"pass"}');
//
//     Http2.send();
//     Http2.onreadystatechange=(e)=>{
//         console.log(Http2.body);
//         console.log("2");
//     };
// };
//
//




(function($){

    console.log("asdfasdf");
    $.ajax({
        url: 'http://127.0.0.1:12345/welcome',
        type: 'get',
        // data: '{"username":"admin", "password":"pass"}' ,
        xhrFields: { withCredentials: true },
        success: function( data, textStatus, jQxhr ){
            console.log( textStatus, jQxhr );
        },
        error: function( jqXhr, textStatus, errorThrown ){
            console.log( errorThrown );
        }
    });
})(jQuery);