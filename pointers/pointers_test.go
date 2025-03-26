package main

import (
	"testing"
)

func TestWallet(t *testing.T){
    assertError:=func(t testing.TB,got error,want error){
        t.Helper()
        if got==nil{
            t.Fatal("got nil wanted an error")
        }
        if got!=want{
            t.Errorf("got %s want %s",got,want);
        }

        
    }
    assertBalance:=func(t testing.TB,w Wallet,want Bitcoin){
        t.Helper()
        got:=w.Balance()
        
        if got!=want{
            t.Errorf("got %s want %s",got,want)
        }
    }
    wallet:= Wallet{}

     wallet.Deposit(Bitcoin(10))
     got:= wallet.Balance()
     want:=Bitcoin(10)
    if got!=want{
        t.Errorf("got %s want %s",got,want);
    }
    t.Run("insufficent funds to withdraw",func(t *testing.T){
        startingBalance:=Bitcoin(10)
        w:=Wallet{startingBalance}
        assertBalance(t,w,startingBalance)
        err:=w.Withdraw(Bitcoin(100))
        assertError(t,err,errInsufficentFunds)


    })
    
}
