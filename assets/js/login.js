$("#login").on("submit", handle_login);

function handle_login(e){
    const pass = $("#password").val();
    const email = $("#email").val();

    e.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email,
            password: pass
        }
    }).done(()=>{
        window.location = "/home";
    }).fail(()=>{
        Swal.fire({
            title: "Ooops...",
            text: "Usuário ou senha inválidos!",
            icon: "error"
        });
    });
}