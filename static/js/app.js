function register() {
    let name = document.getElementById("name").value;
    let email = document.getElementById("email").value;
    let password = document.getElementById("password").value;

    fetch("/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, email, password })
    })
        .then(res => res.json())
        .then(data => {
            if (data.message) {
                alert("Registrazione riuscita! Ora puoi accedere.");
                window.location.href = "/";
            } else {
                alert("Errore: " + data.error);
            }
        });
}

function login() {
    let email = document.getElementById("email").value;
    let password = document.getElementById("password").value;

    fetch("/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password })
    })
        .then(res => res.json())
        .then(data => {
            if (data.token) {
                localStorage.setItem("token", data.token);
                alert("Login riuscito!");
                window.location.href = "/home";
            } else {
                alert("Errore: " + data.error);
            }
        });
}

function logout() {
    localStorage.removeItem("token");
    alert("Logout effettuato!");
    window.location.href = "/";
}
