new Vue({
    el: '#app',
    data:{
        error:"",
        username:"",
        password:""
    },
    methods:{
        signin(){
                axios({ method: "POST", "url": "http:/api/signin", data:'{"username":"' + this.username + '", "password":"' + this.password + '"}' , withCredentials: true}).then(result => {
                window.location.assign('/welcome');
            }, error => {
                this.error="Bad username or password";
            });

        }
    }
});