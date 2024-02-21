#include <iostream>

using namespace std;

/* project 2 c++ lessons */
/* Author Necati Efe Hancer*/
/* ClassNumber 220 723 1011 */

class Product{
	private:
	string productName;
	int productPrice;
    public:
    void setProductName(string name){
    	this->productName=name;
	}
	string getProductName(){
		return productName;
	}
	void setProductPrice(int price){
		if(price<0){
			this->productPrice=0;
		}else{
			this->productPrice=price;
		}
	}
	int getProductPrice(){
		return productPrice;
	}
};


class Customer{
	private:
		string customerName;
		int promosionState;
		int calculatingBill;
		int salingBill;
		Product hameper[30];
	    int arrayIndex;
    public:
    	void setCustomerName(string name){
    		this->customerName=name;
		}
		string getCustomerName(){
			return customerName;
		}
		void setPromosionState(int state){
			if(state<0 && state>100){
				this->promosionState=100;
			}else{
			    this->promosionState=state;
			}
		}
		int getPromosionState(){
			return promosionState;
		}
		void addProduct(Product p){
			if(this->arrayIndex>29){
				cout<<"sepet kapasite dolu 30"<<endl;
			}else{
			this->calculatingBill+=p.getProductPrice();
            this->salingBill += p.getProductPrice() * static_cast<float>(promosionState) / 100.0;
			this->hameper[arrayIndex]=p;
			this->arrayIndex++;
			}
		}
void deleteProduct() {
    string name;
    cout << "urun adini girin ";
    cin >> name;

    int foundIndex = -1;  
    for (int i = 0; i < this->arrayIndex; i++) {
        if (hameper[i].getProductName() == name) {
            foundIndex = i;
            for (int j = i; j < this->arrayIndex - 1; j++) {
                hameper[j] = hameper[j + 1];
            }
            hameper[this->arrayIndex - 1] = Product(); 
            this->arrayIndex--;  
            break;
        }
    }

    if (foundIndex != -1) {
        cout << "urun silindi" << endl;
    } else {
        cout << "urun bulunamadi" << endl;
    }
}

		void listProducts(){
			for(int i=0;i<this->arrayIndex;i++){
				cout<<hameper[i].getProductName()<<"_"<<hameper[i].getProductPrice()<<endl;
			}
			cout<<"toplam tutar"<<calculatingBill<<endl;
			cout<<"indirimli tutar"<<salingBill<<endl;
		}
	Customer(){
		this->customerName="Bilinmiyor";
		this->promosionState=100;
		this->calculatingBill=0;
		this->salingBill=0;
		this->arrayIndex=0;
	}	
};

Product creatingProduct(){
	string name;
	int price;
	cout<<"Malin Adini Giriniz"<<endl;
	cin>>name;
	cout<<"Malin Fiyatini Giriniz"<<endl;
	cin>>price;
	Product np;
	np.setProductName(name);
	np.setProductPrice(price);
	return np;
}

int initilazeCommand(Customer array[], int length) {
    for (int i = 0; i < length; i++) {
        cout << "(" << i+1 << ") " << array[i].getCustomerName() << endl;
    }
    cout<<"(4) Sonlandir"<<endl;
    int flag;
    cin>>flag;
    return flag;
}

int initilazeCommand2(){
  cout<<"(1) Mal Ekle"<<endl;
  cout<<"(2) Mal Sil"<<endl;
  cout<<"(3) Mal Listele"<<endl;
  int flag;
  cin>>flag;
  return flag;
}

int main(int argc, char** argv) {
	
	
Customer customers[3];
Customer efe;
efe.setCustomerName("Necati Efe");
efe.setPromosionState(80);
customers[0]=efe;

Customer ahmet;
ahmet.setCustomerName("Ahmet");
ahmet.setPromosionState(-5);
customers[1]=ahmet;

Customer tacettin;
customers[2]=tacettin;

	
while(true){
    switch(initilazeCommand(customers,3)){
        case 1:
            switch(initilazeCommand2()){
                case 1:
                    customers[0].addProduct(creatingProduct());
                    break;
                case 2:
                    customers[0].deleteProduct();
                    break;
                case 3:
                    customers[0].listProducts();
                    break;
                default:
                    break;
            }
            break; 
        case 2:
            switch(initilazeCommand2()){
                case 1:
                    customers[1].addProduct(creatingProduct());
                    break;
                case 2:
                    customers[1].deleteProduct();
                    break;
                case 3:
                    customers[1].listProducts();
                    break;
                default:
                    break;
            }
            break;
        case 3:
            switch(initilazeCommand2()){
                case 1:
                    customers[2].addProduct(creatingProduct());
                    break;
                case 2:
                    customers[2].deleteProduct();
                    break;
                case 3:
                    customers[2].listProducts();
                    break;
                default:
                    break;
            }
            break; 
        case 4:
            return 0;
            break;
        default:
            cout << "Geçersiz komut" << endl;
            break;
    }
}
	return 0;
}
