$(function () {

	$(".input input").focus(function () {

		$(this).parent(".input").each(function () {
			$("label", this).css({
				"line-height": "18px",
				"font-size": "18px",
				"font-weight": "100",
				"top": "0px"
			})
			$(".spin", this).css({
				"width": "100%"
			})
		});
	}).blur(function () {
		$(".spin").css({
			"width": "0px"
		})
		if ($(this).val() == "") {
			$(this).parent(".input").each(function () {
				$("label", this).css({
					"line-height": "60px",
					"font-size": "24px",
					"font-weight": "300",
					"top": "10px"
				})
			});

		}
	});

	$(".button").click(function (e) {
		var pX = e.pageX,
			pY = e.pageY,
			oX = parseInt($(this).offset().left),
			oY = parseInt($(this).offset().top);

		$(this).append('<span class="click-efect x-' + oX + ' y-' + oY + '" style="margin-left:' + (pX - oX) + 'px;margin-top:' + (pY - oY) + 'px;"></span>')
		$('.x-' + oX + '.y-' + oY + '').animate({
			"width": "500px",
			"height": "500px",
			"top": "-250px",
			"left": "-250px",

		}, 600);
		$("button", this).addClass('active');
	})

	$(".alt-2").click(function () {
		if (!$(this).hasClass('material-button')) {
			$(".shape").css({
				"width": "100%",
				"height": "100%",
				"transform": "rotate(0deg)"
			})

			setTimeout(function () {
				$(".overbox").css({
					"overflow": "initial"
				})
			}, 600)

			$(this).animate({
				"width": "140px",
				"height": "140px"
			}, 500, function () {
				$(".box").removeClass("back");

				$(this).removeClass('active')
			});

			$(".overbox .title").fadeOut(300);
			$(".overbox .input").fadeOut(300);
			$(".overbox .button").fadeOut(300);
			$(".overbox .signup-error").fadeOut(300);

			$(".alt-2").addClass('material-buton');
		}

	})

	$(".material-button").click(function () {

		if ($(this).hasClass('material-button')) {
			setTimeout(function () {
				$(".overbox").css({
					"overflow": "hidden"
				})
				$(".box").addClass("back");
			}, 200)
			$(this).addClass('active').animate({
				"width": "700px",
				"height": "700px"
			});

			setTimeout(function () {
				$(".shape").css({
					"width": "50%",
					"height": "50%",
					"transform": "rotate(45deg)"
				})

				$(".overbox .title").fadeIn(300);
				$(".overbox .input").fadeIn(300);
				$(".overbox .button").fadeIn(300);
				$(".overbox .signup-error").fadeIn(300);
			}, 700)

			$(this).removeClass('material-button');

		}

		if ($(".alt-2").hasClass('material-buton')) {
			$(".alt-2").removeClass('material-buton');
			$(".alt-2").addClass('material-button');
		}

	});
	$("#login").click(function () {
		// alert('login');
	});

	$("#signup").click(function () {
		first_name = $("#regfirstname").val();
		last_name = $("#reglastname").val();
		email = $("#regemail").val();
		pass = $("#regpass").val();

		if (!validateEmail(email)) {
			$("#signup-error").text("Email invalid.");
			return;
		}
		if (first_name == "" || last_name == "" || email == "" || pass == "") {
			$("#signup-error").text("All fields should be filled in.");
			return;
		}

		data = {
			application: {
				client_id: "id",
				client_secret: "secret"
			},
			user: {
				first_name: first_name,
				last_name: last_name,
				email: email,
				password: pass
			}
		}

		$.ajax({
			method: "POST",
			url: "http://localhost:3123/signup",
			crossDomain: true,
			data: JSON.stringify(data),
			contentType: "application/json",
			dataType: "json",
			success: function(data, status) {
				window.location.href = "/";
			},
			error: function(data, status) {
				console.log(data);
				$("#signup-error").text(data.responseJSON.data.errors[0].message);
			}
		});
	});
});

function validateEmail(email) {
	var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
	return re.test(String(email).toLowerCase());
}