function edit(event){
    event.preventDefault();

    $.ajax({
        url: "/profile-update",
        method: "PUT",
        data: {
            name: $("#name").val(),
            email: $("#email").val(),
            nick: $("#nick").val()
        }
    }).done(()=>{
        Swal.fire({
            title: "Sucesso",
            text: "Usuário atualizado com sucesso!",
            icon: "sucess"
        }).then(()=>{
            window.location = "/profile";
        });
    }).fail((err)=>{
        console.log("Error: ", err);
        Swal.fire({
            title: "Ooops...",
            text: "Erro ao atualizar o usuário!",
            icon: "error"
        });
    });;
}

$("#update-form").on("submit", edit);