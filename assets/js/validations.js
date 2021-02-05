$(document).ready(function() { 
 
    $('#btn-submitt').click(function() {  
 
        $(".error").hide();
        var hasError = false;
        var emailReg = /^([\w-\.]+@([\w-]+\.)+[\w-]{2,4})?$/;
       
        var firstname = $("#Userfname").val();
        var lastname = $("#Userlname").val();
        var emailaddressVal = $("#UserEmail").val();        
        var country = $("#country_dropdown").val();        
        var state = $("#state_dropdown").val();        
        var city = $("#city_dropdown").val();        


        if(firstname == '') {
            $("#Userfname").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">*Enter your First name.</span>');
            hasError = true;
          } 

          if(lastname == '') {
            $("#Userlname").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">*Enter your Last name.</span>');
            hasError = true;
          } 


        if(emailaddressVal == '') {
            $("#UserEmail").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">*Enter your email address.</span>');
            hasError = true;
        }
 
        else if(!emailReg.test(emailaddressVal)) {
            $("#UserEmail").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">*Enter a valid email address.</span>');
            hasError = true;
        }
 
       
        if(country == '') {
            $("#country_dropdown").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">* Select country.</span>');
            hasError = true;
          } 

          if(state == '') {
            $("#state_dropdown").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">*Select state.</span>');
            hasError = true;
          } 

          if(city == '') {
            $("#city_dropdown").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">*Select city.</span>');
            hasError = true;
          } 
         
          if(hasError == true) { return false; }



      


    });
});