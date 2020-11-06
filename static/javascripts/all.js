function confimrDelete(id) {
    var r = confirm("Are you sure what to delete?");
    if (r == true) {
      window.location.href='/profiles/delete/'+id;
        return;
    } else {
        return false;
    }
}

function confimrDeleteMod(id,url) {
    var r = confirm("Are you sure what to delete?");
    if (r == true) {
      window.location.href=url+id;
        return;
    } else {
        return false;
    }
}
