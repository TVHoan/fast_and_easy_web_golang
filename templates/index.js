var user;

fetch("/api/userinfor")
    .then(response => response.json())
    .then(data => {
        user = data;
        $("#userinfo").text(data.username) // Process the received data here
    })
    .catch(error => {
        console.log('An error occurred', error);
    });