$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault();

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert("As senhas não coincidem!");
        return;
    }

    $.ajax({
        URL: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            nome: $('#email').val(),
            nome: $('#nick').val(),
            nome: $('#senha').val()
        }
    })
}