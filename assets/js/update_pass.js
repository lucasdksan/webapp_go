function updata_pass(event){
    event.preventDefault();

    if($("#new-pass").val() != $("#check-pass").val()) {
        Swal.fire({
            title: "Ooops...",
            text: "As senhas não coincidem!",
            icon: "warning"
        });

        return;
    }

    $.ajax({
        url: "/update-pass",
        method: "POST",
        data: {
            current: $("#pass").val(),
            newPass: $("#new-pass").val()
        }
    }).done(()=>{
        Swal.fire({
            title: "Sucesso",
            text: "Senha foi atualizada com sucesso!",
            icon: "sucess"
        }).then(()=>{
            window.location = "/profile";
        });
    }).fail((err)=>{
        console.log("Error: ", err);
        Swal.fire({
            title: "Ooops...",
            text: "Erro ao atualizar a senha!",
            icon: "error"
        });
    });
}

$("#update-pass").on("submit", updata_pass);