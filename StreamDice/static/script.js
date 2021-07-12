$(".dice").click(function () {
	let counter = 0;
	$(".dice_count", this).show();
	if ($(".dice_count", this).text()) {
		counter = $(".dice_count", this).text();
		counter = counter - -1;
		$(".dice_count", this).text(counter);
	} else {
		counter = counter - -1;
		$(".dice_count", this).text(counter);
	}
});

$(".roll_btn").click(function () {
	let str = "";
	$(".logs").html("");
	$(".logs").html("ROLL!" + "<br>");
	$('.dice_count').each(function (i, obj) {
		if ($(obj).text() != "") {
			$(".logs").html($(".logs").html() + '>' + $(obj).attr("id") + ': ' + $(obj).text() + "<br>");
			str += $(obj).attr("id")+ ':' + $(obj).text() + ';';
		}		
	});

	$.ajax({
		url: '/ajax',
		method: 'POST',
		data: { sendedData: str },
		dataType: 'json',
	    success: function (response) {
			console.log(response)
			//$(".result").text(obj["d100"]);
		},
		error: function (jqXHR, exception) {
			var msg = '';
			if (jqXHR.status === 0) {
				msg = 'Not connect.\n Verify Network.';
			} else if (jqXHR.status == 404) {
				msg = 'Requested page not found. [404]';
			} else if (jqXHR.status == 500) {
				msg = 'Internal Server Error [500].';
			} else if (exception === 'parsererror') {
				msg = 'Requested JSON parse failed.';
			} else if (exception === 'timeout') {
				msg = 'Time out error.';
			} else if (exception === 'abort') {
				msg = 'Ajax request aborted.';
			} else {
				msg = 'Uncaught Error.\n' + jqXHR.responseText;
			}
			console.log(msg);
		},
	});
});

$(".clear_btn").click(function () {
	$(".dice_count").text("");
	$(".dice_count").hide();
	$(".logs").html("");
});