package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vultisig/airdrop-registry/internal/models"
)

func (a *Api) addCoin(c *gin.Context) {
	var coin models.CoinBase
	if err := c.ShouldBindJSON(&coin); err != nil {
		a.logger.Error(err)
		c.Error(errInvalidRequest)
		return
	}
	ecdsaPublicKey := c.Param("ecdsaPublicKey")
	eddsaPublicKey := c.Param("eddsaPublicKey")
	hexChainCode := c.GetHeader("x-hex-chain-code")
	if hexChainCode == "" {
		c.Error(errForbiddenAccess)
		return
	}
	// Ensure the relevant vault exist
	vault, err := a.s.GetVault(ecdsaPublicKey, eddsaPublicKey)
	if err != nil {
		a.logger.Error(err)
		c.Error(errVaultNotFound)
		return
	}
	if vault.HexChainCode != hexChainCode {
		c.Error(errForbiddenAccess)
		return
	}
	addr, err := vault.GetAddress(coin.Chain)
	if err != nil {
		a.logger.Error(err)
		c.Error(errFailedToGetAddress)
		return
	}
	if coin.Address != addr {
		c.Error(errAddressNotMatch)
		return
	}
	coinDB := models.CoinDBModel{
		CoinBase: coin,
		VaultID:  vault.ID,
	}
	if err := a.s.AddCoin(&coinDB); err != nil {
		a.logger.Error(err)
		c.Error(errFailedToAddCoin)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"coin_id": coinDB.ID})
}

func (a *Api) deleteCoin(c *gin.Context) {
	ecdsaPublicKey := c.Param("ecdsaPublicKey")
	eddsaPublicKey := c.Param("eddsaPublicKey")
	strCoinID := c.Param("coinID")
	hexChainCode := c.GetHeader("x-hex-chain-code")
	if hexChainCode == "" {
		c.Error(errForbiddenAccess)
		return
	}

	// Ensure the relevant vault exist
	vault, err := a.s.GetVault(ecdsaPublicKey, eddsaPublicKey)
	if err != nil {
		a.logger.Error(err)
		c.Error(errVaultNotFound)
		return
	}

	if vault.HexChainCode != hexChainCode {
		c.Error(errForbiddenAccess)
		return
	}

	if err := a.s.DeleteCoin(strCoinID, vault.ID); err != nil {
		a.logger.Error(err)
		c.Error(errFailedToDeleteCoin)
		return
	}
	c.Status(http.StatusNoContent)
}

func (a *Api) getCoin(c *gin.Context) {
	strCoinID := c.Param("coinID")
	coin, err := a.s.GetCoin(strCoinID)
	if err != nil {
		a.logger.Error(err)
		c.Error(errFailedToGetCoin)
		return
	}
	c.JSON(http.StatusOK, coin.CoinBase)
}
