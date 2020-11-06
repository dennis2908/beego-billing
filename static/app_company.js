$(document).ready(function() {
	
	$("#Company_code").focus();

    $.ajaxSetup({
		headers: {
			'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
		}
	});
	
    $.get("/company/view/"+$("#search").val(), function(data) {
        $("#table_content > tbody").html(data);
    });
	
	$("#search").keyup(function(e){
	  $.get("/company/view/"+$(this).val(), function(data) {
        $("#table_content > tbody").html(data);
	  });

	});

    $('#add_save').click(function(e) {
		var errors = 0;
		$("#company_submit :input").map(function(){
			 if( !$(this).val() ) {
				  $(this).parents('td').addClass('warning');
				  errors++;
			} else if ($(this).val()) {
				  $(this).parents('td').removeClass('warning');
			}   
		});
		if(errors > 0){
			alert("All fields are required")
			return false;
		}
	
        $.ajax({
            url: "company/save",
            type: "GET",
            data: $('#company_submit').serialize()+"&Id="+$("#Id").val(),
            success: function(data, status, xhr) {
	
				$.when($.get("company/view", function(html) {
		            $("#table_content > tbody").html(html);
                }).done(function() {
					$.when($('#company_submit').trigger("reset")).done(function ( ) {
					$("#Company_code").attr("readonly",false);	
					$("#Company_code").focus();
					e.preventDefault();
				});
					 
				  
				}));
				
               
            },
        });
    }); // add close
	
	$('.del').click(function() {
		var id = $(this).attr('id');
		$.ajax({
	    url : "company/delete/"+id,
	    type: "POST",
	    success: function(data)
	    {  
	    	$.when($.get("company/view", function(html) {
		            $("#table_content > tbody").html(html);
                }).done(function() {
					$.when($('#company_submit').trigger("reset")).done(function ( ) {
					$("#Company_code").attr("readonly",false);	
					$("#Company_code").focus();
				});
					 
				  
				}));
	    }
		});
	});


});

function edit_company(id){

	$.ajax({
	    url : 'company/edit/'+id,
	    type: 'GET',
	    success: function(data)
	    {
			 $.each(data, function( key, value ) {
				  $("#"+key).val(value);
			});
			
			$("#Company_code").attr("readonly","readonly");
			 
			
	    }
	})
	
}

function del_company(id){

	$.ajax({
	    url : 'company/delete/'+id,
	    type: 'GET',
	    success: function(data)
	    {
			 $.when($.get("company/view", function(html) {
		            $("#table_content > tbody").html(html);
                }).done(function() {
					$.when($('#company_submit').trigger("reset")).done(function ( ) {
					$("#Company_code").attr("readonly",false);	
					$("#Company_code").focus();
				});
					 
				  
			}));
			 
			
	    }
	})
	
}