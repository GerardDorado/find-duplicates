# find-duplicates

The main idea with the program is to find possible duplicates among a list of contacts.

The program will search in a list of contact for any possible duplicates among them. This list of possible matches will have a score based on how accurate is the match, this score goes from 1 to 5 where 1 is a low accuracy and 5 is a high accuracy

The steps that the program would take is:

- Read the contact file, normalize its data (removing whitespaces and making everything lowercase)and store it in a data structure

- Comparing each contact with each other, trying to find matches and a assign a score to the match
.
- Writing the output into a csv file.

To run the program:

    -Be sure that there is an input file names "contact.csv" in the same folder as the program
    -Execute "go run main.go"
    -You should see that an "output.csv" file was created

To run the tests:
    -Execute "go test *.go"
