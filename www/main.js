(function() {
	var editor = ace.edit("code");
	editor.setTheme("ace/theme/tomorrow_night_eighties");
	editor.session.setMode("ace/mode/golang");
	editor.focus();

	$("#code").on("keypress", function(e) {
		var meta = e.metaKey || e.ctrlKey;
		var enterKey = (e.keyCode || e.which);
		var enter = enterKey === 10 || enterKey === 13;
		if (meta && enter) {
			$("#output").html("");
			$("#output").append("<p class='ide'>Executing...</p>");
			$.ajax({
				url: '/api/run',
				method: 'POST',
				data: {
					code: editor.getValue()
				},
				success: function(data) {
					$("#output").append("<p class='msg'>" + data.replace("\n", "<br/>")  + "</p>");			
				}
			});
		}
	});

})();
