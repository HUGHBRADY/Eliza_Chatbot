// ========= Eliza Chatbot Project =========
// ========= Hugh Brady  G00338260 =========

const conversation = $("#output");          // <-- what appears in the output box
const user = $("#user-input");              // <-- what the user entered

// We enter this function whenever user types a character into the form
user.keypress(function(e){
    // If the key entered is not Enter, the function will keep going 
    // Otherwise the function will continue
    if(e.keyCode != 13){
        return;
    }

    e.preventDefault();                     // Stops page from refreshing

    const input = user.val();               // <-- assigns what's typed into the constant input
    user.val(" ");                          // empties the input box after you hit enter

    // displays the user's input in the conversation box
    conversation.append("<li class='list-group-item text-left list-group-item-success'>" + "User: " + input + " </li>");

    // used for passing input into go code
    const queryParams = {"user-input" : input}

    // calls Eliza.go's chat handler 
    $.get("/chat", queryParams)

        .done(function(resp){
            setTimeout(function(){
                // displays the bot's response in the conversation box
                conversation.append("<li class='list-group-item text-left list-group-item-success'>" + "Douglas: " + resp + " </li>");
                // moves the view of the page in line with new entries into the conversation
                $("html, body").scrollTop($("body").height());
            }, 1500); //set timeout for a delay between user input and bot response
        }).fail(function(){
            conversation.append("<li class='list-group-item text-left list-group-item-success'>" + "Douglas: Sorry but there's something wrong with my code" + "</li>");
    });
})