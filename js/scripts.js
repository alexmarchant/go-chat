/** @jsx React.DOM */

var host = document.location.host;
var connection = new WebSocket('ws://' + host + '/ws');

var MessageList = React.createClass({
  render: function() {
    var createMessage = function(message) {
      return (
        <div key={message.id} className="chat-message">
          <span className="chat-author">{message.author}: </span>
          <span className="chat-body">{message.body}</span>
        </div>
      );
    };
    return <div>{this.props.messages.map(createMessage)}</div>;
  }
});

var ChatApp = React.createClass({
  getInitialState: function() {
    var randomId = getRandomInt(1, 1000000);
    return {
      messages: [{id: 0, author: "Server", body: "Welcome!"}],
      authorName: randomId,
      authorId: randomId,
      currentMessageBody: "",
    };
  },
  componentDidMount: function() {
    var self = this;
    connection.onopen = function () {
      self.displayMessage({author: "Server", body: "Connected..."});
    };
    connection.onerror = function (e) {
      self.displayMessage({author: "Server", body: "Problem with connection..."})
    };
    connection.onclose = function (e) {
      self.displayMessage({author: "Server", body: "Disconnected!"})
      console.log(e.code);
      console.log(e.reason);
    };
    connection.onmessage = function (message) {
      try {
        var json = JSON.parse(message.data);
        console.log(json);
        if (json.authorId !== self.state.authorId) {
          self.displayMessage(json);
        }
      } catch (e) {
        console.log('This doesn\'t look like a valid JSON: ', message.data);
        return;
      }
    };
  },
  handleSubmit: function(e) {
    e.preventDefault();
    var newMessage = {
      authorId: this.state.authorId,
      author: this.state.authorName,
      body: this.state.currentMessageBody,
    };
    var newCurrentMessageBody = "";
    this.displayMessage(newMessage);
    this.setState({currentMessageBody: newCurrentMessageBody});
    connection.send(JSON.stringify(newMessage));
  },
  displayMessage: function(message) {
    var newId = this.state.messages[this.state.messages.length-1].id + 1;
    message.id = newId;
    var allMessages = this.state.messages.concat(message);
    this.setState({messages: allMessages});
  },
  currentMessageBodyChange: function(e) {
    this.setState({currentMessageBody: e.target.value});
  },
  authorNameChange: function(e) {
    this.setState({authorName: e.target.value});
  },
  render: function () {
    return (
      <div>
        <h1>Chat</h1>
        <div id="chat-window">
          <div id="chat-content">
            <MessageList messages={this.state.messages} />
          </div>
        </div>
        <form onSubmit={this.handleSubmit}>
          <p>
            <label>Name: </label>
            <input type="text" name="authorName" onChange={this.authorNameChange} value={this.state.authorName} />
          </p>
          <p>
            <label>Message: </label>
            <input type="text" name="message" onChange={this.currentMessageBodyChange} placeholder="Type here!" value={this.state.currentMessageBody} />
          </p>
          <p>
            <input type="submit" value="Submit" />
          </p>
        </form>
      </div>
    );
  }
});

React.renderComponent(<ChatApp />, document.body);

function getRandomInt(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}
