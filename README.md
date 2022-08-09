# HouseHoldAPP

Grocery list

GET /List           //Hente komplet liste
POST /AddItem       //Tilføj et item
Delete /DeleteItem  //Slet et item
PATCH /ChangeItem   //Ændre et givent Item

Item 
    Name string
    Volume int
    Unit string

List
    ListID
    WeekNo int
    HouseholdId string
    Items []Item


Meal planner
    - Tilføj egne retter 
        POST /CreateMeal            //Tilføj et måltid
        Get  /GetMeal               //Hent et måltid
        DELETE /DeleteMeal          //Slet et måltid
        PATCH  /UpdateMeal          //Ændre et måltid
        ########## Future work ########
        POST /ShareMeal             //Del et måltid

Meal {
    Name string
    HouseHold string
    Picture string
    Items []item
}

    - Create en uge plan
        POST /CreatePlan            //Tilføj en plan
        Get  /GetPlan               //Hent en plan
        DELETE /DeletePlan          //Slet en plan
        PATCH  /UpdatePlan          //Ændre en plan

WeekPlan {
    WeekNo int
    WeekPlanId //et id der binder denne weekplan op til brugeren / household hvis denne eksistere. 
    Meals []Meal   
}
    - Auto generate på baggrund af egne retter
         /POST /AutoGenerate         //Autogenerer
            FUNCTIONALITY
            Vælge hvor mange måltider der skal være.
            Skammekrog ikke samme måltid i streg. 

    - Udfyld en handle liste på baggrund af mealplan
        PATCH /GenerateList      //Tilføj indkøbsliste til grocery list funktionen.  


User System
    FirstName                       //String
    LastName                        //String
    Email                           //String, Unique
    Password []byte                 //BCrypt
    Household UUID
        POST /Create                //Create household
        POST /InviteUser            //Inviter på baggrund af email
        DELETE /RemoveUser          //Smid en for porten
        DELETE /DeleteHouseHold     //Slet household
        GET /Household              //Hent household data


UserSystem ---- Server ------ DB     ServerPORT: 5001   dbPORT:  27017

Grocery list ---- backend ------ DB  ServerPORT: 5003   dbPORT:  27019

Mealplanner ----- backend ------ DB  ServerPORT: 5005   dbPORT:  27021









