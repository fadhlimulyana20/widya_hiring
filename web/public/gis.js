function handleCredentialResponse(response) {
    // decodeJwtResponse() is a custom function defined by you
    // to decode the credential response.
    const claims = jose.decodeJwt(response.credential)

    console.log(response)
    console.log(claims)
    // const responsePayload = decodeJwtResponse(response.credential);

    // console.log("ID: " + responsePayload.sub);
    // console.log('Full Name: ' + responsePayload.name);
    // console.log('Given Name: ' + responsePayload.given_name);
    // console.log('Family Name: ' + responsePayload.family_name);
    // console.log("Image URL: " + responsePayload.picture);
    // console.log("Email: " + responsePayload.email);
}