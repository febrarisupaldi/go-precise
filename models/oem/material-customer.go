package oem

import(
	"net/http"
	"github.com/febrarisupaldi/go-precise/db"
	r "github.com/febrarisupaldi/go-precise/models"
)

type MaterialCustomer struct{
	Id int `json:"material_customer_hd_id"`
	PCode string `json:"product_code"`
	PName string `json:"product_name"`
	CCode string `json:"customer_code"`
	CName string `json:"customer_name"`
	Status string `json:"Status aktif"`
	CreatedOn string `json:"created_on"`
	CreatedBy string `json:"created_by"`
	UpdatedOn *string `json:"updated_on"`
	UpdatedBy *string `json:"updated_by"`
}

type DetailMaterialCustomer struct{
	CId int `json:"customer_id"`
	MId int `json:"material_id"`
	IsActive bool `json:"is_active"`
	MaterialCustomer
	Detail interface{} `json:"detail"`
}

func AllMaterialCustomer()(r.Response, error){
	var res r.Response
	var obj MaterialCustomer
	var arrobj []MaterialCustomer

	db.Init()
	con := db.Conn()
	sqlStatement := `select material_customer_hd_id, product_code, product_name, customer_code, customer_name,
						case hd.is_active 
							when 0 then 'Tidak aktif'
							when 1 then 'Aktif' 
						end as 'Status aktif',
						hd.created_on, hd.created_by , hd.updated_on, hd.updated_by
					from precise.material_customer_hd hd
					left join precise.product prod on hd.material_id = prod.product_id
					left join precise.customer cust on hd.customer_id = cust.customer_id`
	defer con.Close()

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.PCode, &obj.PName, &obj.CCode, &obj.CName, &obj.Status, &obj.CreatedOn, &obj.CreatedBy, &obj.UpdatedOn, &obj.UpdatedBy)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = arrobj

	return res, nil
}

func ShowMaterialCustomer(id int)(DetailMaterialCustomer, error){
	type Detail struct{
		Id int	`json:"material_customer_dt_id"`
		HId int	`json:"material_customer_hd_id"`
		PCId int `json:"product_customer_id"`
		PCode string `json:"product_code"`
		PName string `json:"product_name"`
		IsActive bool `json:"is_active"`
		CreatedOn string `json:"created_on"`
		CreatedBy string `json:"created_by"`
		UpdatedOn *string `json:"updated_on"`
		UpdatedBy *string `json:"updated_by"`
	}
	var detail Detail
	var obj DetailMaterialCustomer
	var arrobj []Detail

	db.Init()
	con := db.Conn()
	sqlStatement := `select
						material_customer_hd_id, 
						material_id, product_code, product_name, 
						hd.customer_id, customer_code, customer_name,
						hd.is_active,
						hd.created_on, hd.created_by, 
						hd.updated_on, hd.updated_by
					from precise.material_customer_hd hd
					left join precise.product prod on hd.material_id = prod.product_id
					left join precise.customer cust on hd.customer_id = cust.customer_id
					where material_customer_hd_id = ?`
	err := con.QueryRow(sqlStatement, id).Scan(
		&obj.Id, &obj.MId, &obj.PCode, &obj.PName, &obj.CId, &obj.CCode, &obj.CName,
		&obj.IsActive, &obj.CreatedOn, &obj.CreatedBy, &obj.UpdatedOn, &obj.UpdatedBy,
	)
	defer con.Close()

	if err != nil {
		return obj, err
	}

	sqlStatement = `select
						material_customer_dt_id, material_customer_hd_id,
						dt.product_customer_id, product_code, product_name, 
						dt.is_active, 
						dt.created_on, dt.created_by, 
						dt.updated_on, dt.updated_by
					from precise.material_customer_dt dt
					left join precise.product_customer pc on dt.product_customer_id = pc.product_customer_id
					left join precise.product prod on pc.product_id = prod.product_id
					where material_customer_hd_id = ?`
	
	rows, err := con.Query(sqlStatement, id)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&detail.Id, &detail.HId, &detail.PCId, &detail.PCode, &detail.PName,
			&detail.IsActive, &detail.CreatedOn, &detail.CreatedBy, &detail.UpdatedOn, &detail.UpdatedBy,)

		if err != nil {
			return obj, err
		}

		arrobj = append(arrobj, detail)
	}

	obj.Detail = arrobj



	return obj, nil
}