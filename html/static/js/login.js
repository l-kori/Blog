$('.login-btn').click(function() {
	const username = $('#username').val();
	const password = $('#password').val();
	if (!username || !password) {
		layer.msg('用户名或者密码不允许为空', {
			icon: 5,
			time: 1000,
			anim: 1,
			offset: '50px'
		});
		return;
	}
	const data = {
		username: username,
		password: password
	};
	login(data);
});
$('.register').click(function() {
	window.location.href = 'register.html';
});
