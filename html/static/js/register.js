$('.register-btn').click(function() {
  const username = $('#username').val()
  const password = $('#password').val()
  const confirmPassword = $('#confirm-password').val()

  if (!username || !password || !confirmPassword) {
    layer.msg('用户名或者密码不允许为空', {
      icon: 5,
      time: 1000,
      anim: 1,
      offset: '50px'
    })
    return
  }
  if (password !== confirmPassword) {
    layer.msg('两次输入的密码不同', {
      icon: 5,
      time: 1000,
      anim: 1,
      offset: '50px'
    })
    return
  }
  const data = {
    username: username,
    password: password
  }
  register(data)
})
