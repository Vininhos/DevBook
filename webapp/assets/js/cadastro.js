$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault();

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        Swal.fire("Erro", "As senhas não coincidem!", "error");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        dataType: "text",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done(function () {
        Swal.fire("Sucesso!", "Usuário cadastrado com sucesso.", "success").then(function () {
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email: $('#email').val(),
                    senha: $('#senha').val()
                }
            }).done(function () {
                window.location = "/home";
            }).fail(function () {
                Swal.fire("Erro!", "Erro ao autenticar o usuário.", "error");
            })
        });
    }).fail(function (erro) {
        Swal.fire("Erro", "Erro ao cadastrar o usuário!", "error");
    });
}