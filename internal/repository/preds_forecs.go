package repository

import (
	"fmt"
	"strconv"
	"strings"
	"web/internal/ds"
)

func (r *Repository) Preds_forecsList(pr_id string) (*[]ds.Preds_Forecs, error) {
	var Preds_Forecs []ds.Preds_Forecs
	r.db.Where("prediction_id = ?", pr_id).Find(&Preds_Forecs)
	return &Preds_Forecs, nil
}

func (r *Repository) GetPredForecByID(id1 string, id2 string) (*ds.Preds_Forecs, error) {
	var Preds_Forecs ds.Preds_Forecs
	intId1, _ := strconv.Atoi(id1)
	intId2, _ := strconv.Atoi(id2)
	r.db.Find(&Preds_Forecs, intId1, intId2)
	return &Preds_Forecs, nil
}

func (r *Repository) CreatePreds_Forecs(prediction_id string, forecast_id string) error {
	var n ds.Preds_Forecs
	n.ForecastID, _ = strconv.Atoi(forecast_id)
	n.PredictionID, _ = strconv.Atoi(prediction_id)
	return r.db.Create(&n).Error
}

func (r *Repository) DeletePreds_Forecs(prediction_id string, forecast_id string) {
	query := "DELETE FROM preds_forecs WHERE prediction_id = $1 and forecast_id = $2"
	r.db.Exec(query, prediction_id, forecast_id)
}

func (r *Repository) GetForecastsByID(pred_id string) (*[]ds.Forecs_inputs, error) {
	var prf []ds.Preds_Forecs
	r.db.Where("prediction_id = ?", pred_id).Find(&prf)
	var forecs []ds.Forecs_inputs
	var tmp ds.Forecs_inputs
	for i := range prf {
		f, err := r.GetForecastByID(strconv.Itoa(prf[i].ForecastID))
		if err != nil {
			return nil, err
		}
		tmp.Forecast = *f
		tmp.Input = prf[i].Input
		forecs = append(forecs, tmp)
	}
	return &forecs, nil
}
func (r *Repository) GetPredLen(pred_id string) int {
	var prf []ds.Preds_Forecs
	r.db.Where("prediction_id = ?", pred_id).Find(&prf)
	return len(prf)
}

func (r *Repository) SaveInputs(pr_id int, ids []string, vals []string) {
	var pr_fc ds.Preds_Forecs
	for i := range ids {
		pr_fc.Input = vals[i]
		r.db.Model(&pr_fc).Where("prediction_id = ? and forecast_id = ?", pr_id, ids[i]).Updates(&pr_fc)
		fmt.Println(i, ids[i], vals[i])
	}
}
func (r *Repository) EditPredForec(f_id string, pr_id string, input string) error {
	var pred_forec ds.Preds_Forecs

	if err := r.db.Where("forecast_id = ? AND prediction_id = ?", f_id, pr_id).First(&pred_forec).Error; err != nil {
		return err
	}

	if err := r.db.Save(&pred_forec).Error; err != nil {
		return err
	}

	return nil
}

func Calculate(window int, amount int, input string) ([]int, error) {
	int_arr, err := ValidateInput(input)
	if err != nil {
		return nil, err
	}
	if len(int_arr) < window {
		return nil, fmt.Errorf("invalid window value")
	}
	//PREDICTION CALCULATION
	predictions := make([]int, amount)
	data_len := len(int_arr)
	start, end := 0, 0
	windowSum := 0
	//get first window down
	for end < amount {
		windowSum += int_arr[end]
		end++
	}

	delta_sums := 0                     //to find an average trend between deltas
	windows_count := data_len - end + 1 //amount of windows
	var delta int
	//calculate the average sums and the trend
	for end < data_len {
		delta = int_arr[end] - int_arr[start]
		delta_sums = delta_sums + delta
		windowSum += delta
		start++
		end++
	}
	//anylize the data recieved
	int_arr = append(int_arr, predictions...)
	delta_trend := delta_sums / windows_count
	//~predict the future~
	for end < len(int_arr) {
		windowSum += delta_trend
		int_arr[end] = windowSum - int_arr[start]
		start++
		end++
	}
	return int_arr[data_len:], nil
}
func ValidateInput(input string) ([]int, error) {
	withoutsp := strings.ReplaceAll(input, " ", "")
	splitstr := strings.SplitAfter(withoutsp, ",")
	var result []int
	for i := range splitstr {
		curr, err := strconv.Atoi(splitstr[i])
		if err != nil {
			return nil, fmt.Errorf("invalid input")
		}
		result = append(result, curr)
	}
	return result, nil
}
