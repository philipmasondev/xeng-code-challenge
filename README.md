# Xeng Coding Challenge

#### -- Project Status: Privately working on skill evaluation project.

## Project Intro/Objective
The purpose of this project is to provide baseline a skill proficiency. This project has three objectives.
-Pull data from an API selectively.
-Store that data in persistent database.
-Have API calls to pull selective searches against that database. 

### Technologies
* Golang
* PostGres
* Pandas
* HTML
* Bootstrap

## Needs of this project

- Frontend developers
- data exploration/importation
- data processing/cleaning

## Getting Started

1.	Home page provides a button to pull data from API and store this into persistence PostGres Database. 
a.	Future update to this is planned to have user authentication so that only authenticated users can perform API calls. 
2.	/list returns all records via API.  
3.	/search/?ingredient= returns recipes with ther search term.
    - http://localhost:8080/search/?ingredient=butter
    
