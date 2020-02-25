# encrypted-link-generator
Basic CSV parser that generates encrypted PassKit Smart Links from a CSV, which can then be distributed to your customers via your preferred channels. Takes an input file, output destination, base URL and encryption key.

Generating links is free. A pass record is only created once someone clicks on the link; so this is a very
economical way to generate all your unique pass URLs, that already contain the encrypted pass data, and distribute them to your customers.

The tool is quite fast - it generated over 600,000 links in less than 7 seconds on a Dual Core MacBook Pro with 16GB RAM.

## Important: 
* please note that the generated links will only work if you project is set public (figure 1). For private projects you can only create passes via the API or Admin Portal.
* this tool is for the latest PassKit v4 platform: https://app.passkit.com. The tool does not working with older versions of PassKit (Cherry Pie, and the v2/v3 API's).

## Prerequisites:
* a PassKit account. Sign up for free at: https://app.passkit.com
* a project setup in your PassKit account.

## How to use:
* Either compile (`go build src/main.go`) yourself or download any of the pre-compiled executables from the `bin` folder (you may need to `chmod 755` to run it).
* Get your key for your project from the PassKit portal: https://app.passkit.com. You can find the key in your Project Distribution Settings page (figure 2).
* Get your landing page url from your Project Distribution Settings page (figure 3).
* Create your CSV file (you can find a sample in the examples folder).
* Ensure the CSV file contains the correct [headers](#usable-field-names), otherwise we won't be able to map it to the correct pass field.
* Run the tool: `/bin/encrypted-link-generator-osx  -inFile examples/example-data.csv -outFile out.csv -baseURL https://pskt.io/c/wrsynr -encryptionKey f33332e108e3e5e040924d7dd7651f6f54b242525bf4e8733ea12ac3538af755` (replace values with your own)
* Open up out.csv and see the links added against each record.
* You can now distribute the links to your customers. 

## Images
Figure 1:
![Figure 1](https://passkit.com/images/github/passkit-public-setting.png "PassKit Project Settings")

Figure 2:
![Figure 2](https://passkit.com/images/github/passkit-key.png "PassKit Distribution Settings - Project Key")

Figure 3:
![Figure 3](https://passkit.com/images/github/passkit-project-url.png "PassKit Distribution Settings - Project URL")


## Command Parameters
* `-inFile`: The path to your CSV with data.
* `-outFile`: The filename that the program creates & outputs to (if the file exists it will be overwritten).
* `-baseURL`: Your project URL (can get from the PassKit Portal).
* `-encryptionKey`: Your projects encryption key (can get from the PassKit Portal).

## Available Field Names
The following fields are currently supported. These are the field names to use as headers in your CSV.

### Person Fields
* person.surname
* person.forename
* person.otherNames
* person.salutation
* person.suffix
* person.displayName
* person.gender
* person.dateOfBirth
* person.emailAddress
* person.mobileNumber

### Member Specific Fields
* members.tier.name
* members.member.externalId
* members.member.points
* members.member.tierPoints
* members.member.secondaryPoints
* members.member.groupingIdentifier
* members.member.profileImage

### Generic Fields
* universal.optInt

### UTM Tracking Parameters
* utm_source
* utm_medium
* utm_name
* utm_term
* utm_content

## Support
The CSV tool works with all published PassKit protocols: members & coupons (soon to be released). 
