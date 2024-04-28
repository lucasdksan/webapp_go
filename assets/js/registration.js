$("#formulario-cadastro").on("submit", create_user);

function create_user(e){
    e.preventDefault();
    const pass = $("#password").val();
    const check_pass = $("#check-password").val();
    const name = $("#name").val();
    const email = $("#email").val();
    const nick = $("#nick").val();

    if(pass !== check_pass) {
        alert("As senhas são diferentes!");
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name,
            email,
            nick,
            password: pass
        }
    }).done(()=>{
        alert("Usuário cadastro com sucesso!");
    }).fail((err)=>{
        console.log("Error: ", err);
        alert("Erro ao cadastrar o usuário!");
    });
}