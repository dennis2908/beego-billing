<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>Company</h1>
     <p class="lead">Here you'll find all the available company</p>
   </div>
 </div>
</div>
</div>
<div class="container">
        <div class="row">
            <div class="col-md-12">
                <form id="company_submit">

						  <div class="form-group">
							<label for="Company_code">Company Code</label>
							<input type="text" class="form-control" id="Company_code" style="width:500px"  name="Company_code" placeholder="Enter Company Code" required>
						  </div>
						  <div class="form-group">
							<label for="Company_name">Company Name</label>
							<input type="text" class="form-control" id="Company_name" style="width:500px"  name="Company_name" placeholder="Enter Company Name" required>
						  </div>
						  <div class="form-group">
							<label for="Phone_number">Phone Number</label>
							<input type="text" class="form-control" id="Phone_number" style="width:500px"  name="Phone_number" placeholder="Enter Phone Number" required>
						  </div>
						  <div class="form-group">
							<label for="Address">Address</label>
							<textarea class="form-control" id="Address" name="Address" style="width:500px" placeholder="Enter Address" required></textarea>
						  </div>
						
			   </form>
			   <input type="hidden" id="Id" >
			   <button type="submit" id="add_save" class="btn btn-primary">Submit</button>
			   
            </div>
        </div>
        <div class="row">
            <div class="col-md-12">
                <div id="records_content"></div>
                <br>
                <div class="col-md-offset-1 col-md-10" id="table_content">
                </div>
            </div>
        </div>
    </div>
	<div class="row" style="text-align:right">
		  <div class="col-md-6">
			 <input class="form-control" id="search" type="text" placeholder="Search Company Code Or Company Name Here" style="margin-left:370px" aria-label="Search">
		  </div> 
		</div>
<div class="container">
  <div class="row" >
    <table class="table" id="table_content">
      <thead>
        <tr>
		  <th>No</th>
		  <th>Action</th>	
          <th>Company Code</th>
		  <th>Company Name</th>
		  <th>Address</th>	
          <th>Phone Number</th>
        </tr>
      </thead>
      <tbody>
      </tbody>
    </table>
  </div>
</div>

<script src="/static/js/jquery-2.2.0.min.js"></script>
	<script src="/static/app_company.js"></script>