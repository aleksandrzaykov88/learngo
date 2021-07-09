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
		url: '/',
		method: 'POST',
		data: { sendedData: str },
	});
});

$(".clear_btn").click(function () {
	$(".dice_count").text("");
	$(".dice_count").hide();
	$(".logs").html("");
});