new Vue({
    el: '#app',
    data:{
        error:"",
        username:"",
        password:""
    },
    methods:{
        signup(){
                console.log("asdfasdf");
            axios({ method: "POST", "url": "http://127.0.0.1:12345/api/signup", data:'{"username":"' + this.username + '", "password":"' + this.password + '"}' , withCredentials: true}).then(result => {
                // this.jsonInfo = result.data;
                // var obj = JSON.parse(result.data);
                // console.error(result.data);
                // window.location.assign('/');
                this.error = result.data;
            }, error => {
                console.error(error);
            });
        }
    }
});