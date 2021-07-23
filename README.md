# encrypted-link-generator
Basic CSV parser that generates encrypted PassKit SmartPass Links from a CSV file. Links are distributed to your customers via your preferred channels. 

**Generation of PassKit SmartPass links is completely free - you can generate as many links as you need**.

A PassKit pass record is only created once someone visits the link for the first time. 

This is THE fastest, most economical, and secure way to generate and distribute your 
pass links at scale without the need for Developer Resources or the need to implement the PassKit API.

## Table Of Contents
- [Requirements](#requirements)
- [Important Notes](#important-notes)
- [How To Use](#how-to-use)
- [Performance](#performance)
- [Command Parameters](#command-parameters)
- [Available Field Names](#available-field-names)
- [Examples](#examples)
- [PassKit Portal](#passkit-portal)
- [Getting Help](#getting-help)
- [Contributing](#contributing)
- [Author & License](#author--license)

## Requirements:
* PassKit account: sign up and develop for free at: https://app.passkit.com/signup.
* A project setup in your PassKit account.

## Important Notes: 
* Supports the latest PassKit IO Platform: https://app.passkit.com.
* Does not support older versions of PassKit (Cherry Pie, and the v2/v3 API's).

## How to use:
* Download any of the pre-compiled executables from the `bin` folder for your Operating System (you may need to set the right system permissions to run it).
* Get your Project URL: copy from Distribution >> SmartPass Settings in PassKit IO Portal:

![Project URL](/images/project-url.png)

* Get your Project Encryption Key: copy from Distribution >> SmartPass Settings in PassKit IO Portal:

![Project Encryption Key](/images/project-key.png)

* Create your CSV file (you can find samples in the [examples](/examples) folder; just delete the columns that you don't need).
* Ensure the CSV file contains the correct [headers](#available-field-names) (otherwise we won't be able to map it to the correct pass field).
* Run the tool from the command-line: `./encrypted-link-generator-osx -in in.csv -out out.csv -url https://pskt.io/c/wrsynr -key f33332e108e3e5e040924d7dd7651f6f54b242525bf4e8733ea12ac3538af755` (replace values with your own).
* Open up `out.csv`, test and distribute your links.

## Performance
The tool generates 1M pass links in under 10 seconds (on a MacBook Pro DualCore 16GB RAM).

## Command Parameters
* `-in`: the path to your CSV file.
* `-out`: the file that the tool outputs to (if the file exists it will be overwritten).
* `-url`: your Project URL: copy from Distribution >> SmartPass Settings in PassKit IO Portal.
* `-key`: your Project encryption key: copy from Distribution >> SmartPass Settings in PassKit IO Portal.

## Available Field Names
The following fields are currently supported. These are the field names to use as headers in your CSV. For additional details check the [documentation](https://docs.passkit.io/).

**You only need to provide the fields that are applicable to your project. All fields are optional, unless set as required in project.**

### Generic Fields
* `universal.optIn`: true -or false.
* `universal.expiryDate`: expiry date in valid ISO-8601 format (2020-06-22T00:00:00-05:00).

### Person Fields
* `person.surname`: surname / family name.
* `person.forename`: forename / given name.
* `person.otherNames`: other names.
* `person.salutation`: salutation or title.
* `person.suffix`: suffix. For multiple suffixes, separate with spaces.
* `person.displayName`: if required, a string representing the user's preferred designation.
* `person.gender`: gender, as per government issued id. Possible values: `0` (Not known), `1` (Male) or `2` (Female)
* `person.dateOfBirth`: valid date in yyyy-mm-dd format (1980-06-22).
* `person.emailAddress`: valid email address.
* `person.mobileNumber`: valid telephone number.
* `person.externalId`: external ID for the person. Used for some of the native PassKit integrations.

### Member Specific Fields
* `members.tier.id`: the tier ID to enrol the member into.
* `members.member.externalId`: sets the 'external' ID of the member (i.e. the member ID as it's being used in your system). If provided then this can be used to query & update members. This field will be treated as unique within the program, and cannot be updated at a later stage.
* `members.member.points`: primary points balance of the member.
* `members.member.tierPoints`: tier points for the member.
* `members.member.secondaryPoints`: secondary points balance of the member.
* `members.member.groupingIdentifier`: used to group members under the same membership (i.e. couple).
* `members.member.profileImage`: for Membership Programs that require a profile image on the pass. Can either be an image URL or base64 image string.

### Coupon Specific Fields
* `singleUseCoupons.coupon.externalId`: sets the 'external' ID of the coupon (i.e. the unique coupon code as it's being used in your system). If provided then this can be used to query & update coupon. This field will be treated as unique within the campaign, and cannot be updated at a later stage.
* `singleUseCoupons.coupon.sku`: sku of the coupon. Can be used in the barcode by setting ${singleUseCoupons.coupon.sku} in the Pass Template Design barcode settings.

### Meta Fields
* `meta.*`: any custom fields you have defined in your project.

### UTM Tracking Parameters
* `utm_source`: used to identify where the request is coming from. Defaults to Unknown.
* `utm_medium`: used to identify a medium such as email, app, or cost-per-click advertising.
* `utm_name`: used for keyword analysis. Use campaign to identify a specific product promotion or strategic campaign.
* `utm_term`: used for paid search. Use term to note the keywords for the ad that led to the pass.
* `utm_content`: used for A/B testing and content-targeted ads. Use content to differentiate ads or links that point to the same URL.

### Location Parameters
Used to embed locations on a pass level. Each unique pass can have upto 10 different locations embedded in it.

_LOCATION_NUMBER needs to be 1-10_:

* `location[LOCATION_NUMBER].lat`: a valid float for the location latitude, for example 51.507351.
* `location[LOCATION_NUMBER].lon`: a valid float for the location longitude, for example 1.507351.
* `location[LOCATION_NUMBER].lockScreenMessage`: the message to trigger on the lock-screen when someone is within range of the lat/lon.
* `location[LOCATION_NUMBER].alt`: a valid integer for the location altitude, for example 647 (optional).
* `location[LOCATION_NUMBER].name`: an internal name for the location in the PassKit database (optional).
* `location[LOCATION_NUMBER].position`: the position of the location in the locations array (optional).

_When using the above, lat, lon & lock-screen message are mandatory: all 3 need to be provided for each location._

### Beacon Parameters
Used to embed beacons on a pass level (Apple Wallet only). Each unique pass can have upto 10 different beacons embedded in it.

_BEACON_NUMBER needs to be 1-10_:

* `beacon[BEACON_NUMBER].uuid`: a valid UUID string for the beacon.
* `beacon[BEACON_NUMBER].major`: a valid integer for beacon major. Range between 0-65535.
* `beacon[BEACON_NUMBER].minor`: a valid integer for beacon minor. Range between 0-65535.
* `beacon[BEACON_NUMBER].lockScreenMessage`: the message to trigger on the lock-screen when someone is within range of the beacon.
* `beacon[BEACON_NUMBER].name`: an internal name for the beacon in the PassKit database (optional).
* `beacon[BEACON_NUMBER].position`: the position of the beacon in the beacons array (optional).

_When using the above, uuid, major, minor & lock-screen message are mandatory: all 4 need to be provided for each beacon._

### Color Parameters
Used to override colors on a pass level.

* `colors.backgroundColor`: pass background color.
* `colors.labelColor`: label text color (Apple Wallet only).
* `colors.textColor`: value text color (Apple Wallet only).
* `colors.stripColor`: text over Strip Image color (Apple Wallet only).

### Image Parameters
Used to override images on a pass level. Need to contain a valid [https://docs.passkit.io/common/images/#operation/createImages](PassKit 22 character image ID).

* `images.icon`
* `images.logo`
* `images.appleLogo`
* `images.hero`
* `images.eventStrip`
* `images.strip`
* `images.thumbnail`
* `images.background`
* `images.footer`
* `images.security`
* `images.thumbnail`
* `images.privilige`
* `images.airlineAlliance`
* `images.personalization`
* `images.banner`
* `images.message`
* `images.profile`
* `images.appImage`

## Examples
Example CSV files with all supported field names are found in the [examples](/examples) folder in this repo; just delete the columns that you don't need.

## PassKit Portal
The [https://app.passkit.com](https://app.passkit.com) allows you to easily design loyalty cards, membership cards and coupons for both Apple Wallet and Google Pay. 

Additionally, the PassKit portal facilitates management, distribution and simple analysis of your Mobile Wallet projects.

Best Practices:
- Use the web portal for initial account and project setup.
- Then use the Integration Tools / SDKs / APIs to issue, update and delete your individual passes.

## Getting Help
- [Hashed SmartPass Links](https://help.passkit.com/en/articles/3742778-hashed-smartpass-links)
- [Official PassKit Documentation](https://docs.passkit.io/)
- [support@passkit.com](mailto:support@passkit.com)
- [Online Chat Support](https://app.passkit.com/)

## Contributing
Send bug reports, feature requests and code contributions into this repository.

## Author & License
PassKit Inc.: [support@passkit.com](mailto:support@passkit.com)

Distributed under MIT License. Details available in [license file](LICENSE).