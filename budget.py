class Budget:

    def __init__(self,**category):
        self.category = category
        print(self.category)

    def deposit(self,amount,category):
        balance = (self.category[category])
        balance += amount
        print(f"You deposited the sum of #{amount} to {category}")

    def withdraw(self,amount,category):
        balance = (self.category[category])
        if amount > balance:
            print(f"you have exceeded your balance for {category}")
        else:
            balance -= amount
            print(f"You withdrew the sum of #{amount} from {category}")

    def transfer(self,amount,debit_transfer,credit_transfer):
        debit_category = self.category[debit_transfer]
        credit_category = self.category[credit_transfer]
        
        if amount > debit_category:
            print("Your balance is insufficient for this transfer")
        else:
            debit_category -= amount
            credit_category += amount
            print(f"You have transfered the sum of #{amount} from {debit_transfer} to {credit_transfer}")

    def checkBalance(self,category):
        balance = (self.category[category])
        print(f"Your {category} budget balance is {balance}")

personal_budget = Budget(food=1000, clothing=2000,entertainment=500)

personal_budget.deposit(800,'food')

#check when the balance is sufficient or not
personal_budget.withdraw(500,'clothing')
personal_budget.withdraw(3000,'clothing') 

#check whether amount to transfer does not exceed the budget balance for the category
personal_budget.transfer(200,'entertainment','clothing')
personal_budget.transfer(700,'entertainment','clothing')

#checks balance for all category
personal_budget.checkBalance('food')
personal_budget.checkBalance('clothing')
personal_budget.checkBalance('entertainment')
