new Vue({
    el: '#app',
    data:{
        error:"",
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
            axios({ method: "POST", "url": "http://127.0.0.1:12345/api/signin", data:'{"username":"' + this.username + '", "password":"' + this.password + '"}' , withCredentials: true}).then(result => {
                // this.jsonInfo = result.data;
                // var obj = JSON.parse(result.data);
                // console.error(result.data);
                window.location.assign('/welcome');
                console.error(result.data);
            }, error => {
                this.error="Bad username or password";
                // console.error(error);
            });

        }
    }
});