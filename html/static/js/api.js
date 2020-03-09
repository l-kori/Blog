//基础地址
const BASE_URL = 'http://129.204.49.126:9090/';
//注册地址
const REGISTER_API = 'registerPost';
//登录地址
const LOGIN_API = 'loginpost';

//注册
function register(data) {
	var index = layer.load();

	$.ajax({
		url: BASE_URL + REGISTER_API,
		data: data,
		type: 'post',
		dataType: 'json',
		success: function(res) {
			layer.close(index);
			if (res.code === 200) {
				layer.msg(res.message, {
					icon: 6,
					time: 2000,
					anim: 1,
					offset: '50px'
				});
				setTimeout(function() {
					window.location.href = 'login.html';
				}, 2000);
			} else {
				layer.msg(res.code, {
					icon: 5,
					time: 1000,
					anim: 1,
					offset: '50px'
				});
			}
		}
	});
}

//登录
function login(data) {
	var index = layer.load();

	$.ajax({
		url: BASE_URL + LOGIN_API,
		data: data,
		type: 'post',
		dataType: 'json',
		success: function(res) {
			console.log(res);
			layer.close(index);
			if (res.code === 200) {
				layer.msg(res.message, {
					icon: 6,
					time: 2000,
					anim: 1,
					offset: '50px'
				});
				setTimeout(function() {
					window.location.href = 'index.html';
				}, 1000);
			} else {
				layer.msg(res.code, {
					icon: 5,
					time: 1000,
					anim: 1,
					offset: '50px'
				});
			}
		}
	});
}
