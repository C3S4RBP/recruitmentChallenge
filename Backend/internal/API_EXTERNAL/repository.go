package API_EXTERNAL

import (
	"Backend/db"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StockRepository struct {
	db *gorm.DB
}

func NewStockRepository() *StockRepository {
	return &StockRepository{
		db: db.DB,
	}
}

// para procesos granulaes se puede usar este metodo
func (r *StockRepository) BatchCreateOrUpdateStocks(stocks []Stock) error {
	if len(stocks) == 0 {
		return errors.New("la lista de stocks está vacía")
	}

	tickers := make([]string, len(stocks))
	for i, stock := range stocks {
		tickers[i] = stock.Ticker
	}

	var existingStocks []Stock
	if err := r.db.Where("ticker IN ?", tickers).Find(&existingStocks).Error; err != nil {
		return err
	}

	existingMap := make(map[string]Stock)
	for _, stock := range existingStocks {
		existingMap[stock.Ticker] = stock
	}

	var toCreate []Stock
	var toUpdate []Stock

	for _, stock := range stocks {
		if existingStock, exists := existingMap[stock.Ticker]; exists {
			stock.CreatedAt = existingStock.CreatedAt
			stock.UpdatedAt = time.Now()
			toUpdate = append(toUpdate, stock)
		} else {
			stock.CreatedAt = time.Now()
			stock.UpdatedAt = time.Now()
			toCreate = append(toCreate, stock)
		}
	}

	// Ejecutar operaciones en lotes
	if len(toCreate) > 0 {
		batchSize := 100
		for i := 0; i < len(toCreate); i += batchSize {
			end := i + batchSize
			if end > len(toCreate) {
				end = len(toCreate)
			}
			batch := toCreate[i:end]
			if err := r.db.CreateInBatches(batch, len(batch)).Error; err != nil {
				return err
			}
		}
	}

	if len(toUpdate) > 0 {
		batchSize := 100
		for i := 0; i < len(toUpdate); i += batchSize {
			end := i + batchSize
			if end > len(toUpdate) {
				end = len(toUpdate)
			}
			batch := toUpdate[i:end]
			for _, stock := range batch {
				if err := r.db.Save(&stock).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// para procesos masivos se puede usar este metodousando UPSERT de PostgreSQL
func (r *StockRepository) BatchUpsertStocks(stocks []Stock) error {
	if len(stocks) == 0 {
		return errors.New("la lista de stocks está vacía")
	}

	batchSize := 500
	for i := 0; i < len(stocks); i += batchSize {
		end := i + batchSize
		if end > len(stocks) {
			end = len(stocks)
		}
		batch := stocks[i:end]
		if err := r.db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "ticker"}},
			DoUpdates: clause.AssignmentColumns([]string{"target_from", "target_to", "company", "action", "brokerage", "rating_from", "rating_to", "time", "updated_at"}),
		}).CreateInBatches(batch, len(batch)).Error; err != nil {
			return err
		}
	}

	return nil
}

// GetAllStocks obtiene todos los stocks de la base de datos
func (r *StockRepository) GetAllStocks(limit, offset int) ([]Stock, error) {
	var stocks []Stock
	query := r.db.Order("ticker ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	err := query.Find(&stocks).Error
	return stocks, err
}
