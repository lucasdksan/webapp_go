$("#formulario-cadastro").on("submit", create_user);

function create_user(e){
    e.preventDefault();
    const pass = $("#password").val();
    const check_pass = $("#check-password").val();
    const name = $("#name").val();
    const email = $("#email").val();
    const nick = $("#nick").val();

    if(pass !== check_pass) {
        Swal.fire({
            title: "Ooops...",
            text: "As senhas são diferentes!",
            icon: "error"
        });

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
        Swal.fire({
            title: "Sucesso",
            text: "Usuário cadastro com sucesso!",
            icon: "sucess"
        }).then(()=>{
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email,
                    password: pass
                }
            }).done(()=>{
                window.location = "/home";
            });
        });
    }).fail((err)=>{
        console.log("Error: ", err);
        Swal.fire({
            title: "Ooops...",
            text: "Erro ao cadastrar o usuário!",
            icon: "error"
        });
    });
}