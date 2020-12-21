
function login() {
    var username = document.getElementById("username");
    var password = document.getElementById("password");

    $.ajax({
        method: "POST",
        url: "/login", 
        data: {
            username: username.value,
            password: password.value
        },
        success: function(data) {
            location.reload();
        },
        error: function(response) {
            var username = document.getElementById("username");
            var password = document.getElementById("password");
            password.value = "";

            var usernameValue = username.value;
            var usernameStyle = username.style;

            username.value = response.responseJSON.message;
            username.style.color = "#d44";

            setTimeout(
                function() {
                    username.style = usernameStyle;
                    username.value = usernameValue;
                }.bind(username, usernameValue, usernameStyle),
                500
            );
        }
    });
}
