new Vue({
    el: '#app',
    data:{
        title :"asdfg",
        text1: "",
        username:"",
        password:""
    },
    methods:{
        changeText(){
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
        },
        signin(){
            // username = this.username;
            // password = this.password;
            // (function($){

                console.log("asdfasdf");
                $.ajax({
                    url: 'http://127.0.0.1:12345/api/signin',
                    type: 'post',
                    data: '{"username":"' + this.username + '", "password":"' + this.password + '"}' ,
                    async: false,
                    xhrFields: { withCredentials: true },
                    success: function( data, textStatus, jQxhr ){
                        console.log( textStatus, jQxhr );
                        location.href = '/welcome';
                    },
                    error: function( jqXhr, textStatus, errorThrown ){
                        console.log( errorThrown );
                    }
                });
            // })(jQuery);
        }
    }
});