(function() {
	var editor = ace.edit('code');
	editor.setTheme('ace/theme/tomorrow_night_eighties');
	editor.session.setMode('ace/mode/golang');
	editor.focus();

	var localStorage = window.localStorage || { setItem: function() { }, getItem: function() { return ''; } };
	
	var savedCode = localStorage.getItem('code') || '';
	if (savedCode === '') {
		savedCode = 'import "log"\nlog.Print("Welcome to playgo!")\n';
	}
	editor.setValue(savedCode, 1);

	$('#code').on('keydown', function(e) {
		var meta = e.metaKey || e.ctrlKey;
		var keyCode = (e.keyCode || e.which);
		var enter = keyCode === 10 || keyCode === 13;
		var format = keyCode === 32;
    
    if (meta && format) {
      e.preventDefault();
      $('#output').html('<p class="ide">Formatting...</p>');
      $.ajax({
        url: '/api/format',
        method: 'POST',
        data: {
          code: editor.getValue()
        },
        success: function(data) {
          editor.setValue(data, 1);
          $('#output').html('');
        },
        error: function(xhr, status, text) {
        	var response = xhr.responseText.replace(/\n/g, '<br/>');
					if (response) {
						$('#output').html('<p class="msg-err">' + response + '</p>');			
						localStorage.setItem('code', editor.getValue());
					} else {
						$('#output').html('<p class="msg-err">Looks like the server is not reachable.</p>');			
					}
        }
      });
    }
    
    if (meta && enter) {
			e.preventDefault();
			$('#output').html('');
			$('#output').append('<p class="ide">Executing...</p>');
			$.ajax({
				url: '/api/run',
				method: 'POST',
				data: {
					code: editor.getValue()
				},
				success: function(data) {
					var output = data.replace(/\n/g, '<br/>');
					$('#output').html('<p class="msg">' + output + '</p>');
					localStorage.setItem('code', editor.getValue());
				},
				error: function(xhr, status, text) {
					var response = xhr.responseText.replace(/\n/g, '<br/>');
					if (response) {
						$('#output').html('<p class="msg-err">' + response + '</p>');			
						localStorage.setItem('code', editor.getValue());
					} else {
						$('#output').html('<p class="msg-err">Looks like the server is not reachable.</p>');			
					}
				}
			});
		}
	});

})();
