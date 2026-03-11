<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>Company's Bill</h1>
     <p class="lead">Here you'll find all the Billing charge on company</p>
   </div>
 </div>
</div>
</div>
<div class="container">
        <div class="row">
            <div class="col-md-12">
	            <form id="company_submit">
			              <div class="form-group">
							<label for="Company_code">Company Name</label>
							   <select class="form-control" id="Company_code" style="width:500px" name="Company_code" aria-describedby="Company_code" required>
								{{range $val := .records}}
								<option value="{{$val.Company_code}}">{{$val.Company_code}} - {{$val.Company_name}}</option>
								{{end}}
							   </select>
						  </div>
						   <div class="form-group">
							<label for="Bill">Bill</label>
							<input type="number" class="form-control" id="Bill" style="width:500px"  name="Bill" placeholder="Enter Bill" required>
						  </div>
						  <div class="form-group">
							<label for="Note">Note</label>
							<textarea class="form-control" id="Note" name="Note" style="width:500px" placeholder="Enter Note" required></textarea>
						  </div>
						
			   </form>
			   <div class="form-group">
				   <input type="hidden" id="Id" name="Id">
				   <button type="submit" onclick="add_save()" class="btn btn-primary">Submit</button>
			   </div>
			   <div class="form-group row"  style="margin-left:350px">
					<div class="col-sm-7">
					  <input class="form-control" id="search" type="text" placeholder="Enter Company Code Or Company Name">
					</div>
					<button type="submit" id="search_btn" class="btn btn-primary">Search</button>
				</div>
            </div>
        </div>
</div>
<div class="container">
  <div class="row" >
    <table class="table" id="table_content">
      <thead>
        <tr>
		  <th>No</th>
		  <th>Action</th>	
          <th>Company Code - Company Name</th>
		  <th>Note</th>	
          <th>Bill</th>
        </tr>
      </thead>
      <tbody>
      </tbody>
    </table>
  </div>
</div>


		<script src="/static/js/jquery-2.2.0.min.js"></script>
<script type="text/javascript" src="/static/app_bill.js"></script>