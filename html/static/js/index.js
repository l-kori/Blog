$(function() {
	$.ajax({
		type: 'post',
		url: 'http://129.204.49.126:9090',
		dataType: 'json',
		success: function(info) {
			console.log(info.data[3][2]);
			var htmlStr = template('indexList', { list: info.data });
			$('.main-list').html(htmlStr);
		}
	});
});
