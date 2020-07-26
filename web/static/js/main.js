
function ShowSignup() {
  $(".login-btn").hide();
  $(".email-input").show();
    $(".register-btn").show();
}

 function ShowLogin() {
$(".email-input").hide();
  $(".register-btn").hide();
  $(".login-btn").show();
}

function Login() {
var username =  $(".username").val();
var password = $(".password").val();
var loginData = {"username":username,"password":password} ;
if (!Validation(loginData)) {
  return
}
$.ajax({
  type:"POST",
  url:"/login" ,
  contentType:'application/x-www-form-urlencoded; charset=UTF-8',
  data:loginData,
  success: function(html){
    if ( !html["Status"] ){
      var ErrorMsg = html["Error"];
      var errDiv = `<div class="errmsg">${ErrorMsg}</div>` ;
      $('.login-btn').append(errDiv); 
      return
    }
    var successMsg = html["Status"];
    var errDiv = `<div class="successmsg">${successMsg}</div>` ;
    $('.login-btn').append(errDiv); 


  },
  onerror:function(html){
    alert(html);
  }
});
}

function Register() {
  var username =  $(".username").val();
  var password = $(".password").val();
  var email = $(".email").val();
  var registerData ={"username":username,"password":password,"email":email} ;
  if (!Validation(registerData)) {
    return
  }
  $.ajax({
    type:"POST",
    url:"/register" ,
    contentType:'application/x-www-form-urlencoded; charset=UTF-8',
    data: registerData,
    success: function(html){
      if ( !html["Status"] ){
        var ErrorMsg = html["Error"];
        var errDiv = `<div class="errmsg">${ErrorMsg}</div>` ;
        $('.register-btn').append(errDiv); 
        return
      }
      var successMsg = html["Status"];
      var errDiv = `<div class="successmsg">${successMsg}</div>` ;
      $('.register-btn').append(errDiv); 
  
  
    },
    onerror:function(html){
      alert(html);
    }
  });
  }


function Validation(data) {
if (!data["username"]) {
  var errDiv = `<div class="errmsg"> username cannot be empty </div>` ;
  $('.username').append(errDiv); 
  return  false
}
if (!data["password"]) {
  var errDiv = `<div class="errmsg"> password cannot be empty </div>` ;
  $('.password').append(errDiv); 
  return  false
}
if (data["email"]) {
  if (data["email"]== undefined) {
  var errDiv = `<div class="errmsg"> email cannot be empty </div>` ;
  $('.email').append(errDiv); 
  return  false}
}
return true
}
