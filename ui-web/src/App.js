import React, { useContext, useEffect, useState } from "react";
import { BrowserRouter as Router, Link, Redirect, useHistory, useLocation } from "react-router-dom";
import { Route } from "react-router-dom";
import { v4 as uuidv4 } from 'uuid';

import axios from "axios";
import Provider from "./Provider";
import UserContext from "./UserContext";

const instance = axios.create({
  baseURL: "http://localhost:8080/app",
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    "Access-Control-Allow-Origin": "*",
  },
});

function App() {
  return (
    <Router>
      <Provider>
        <BaseLayout/>
      </Provider>
    </Router>
  );
}

const BaseLayout = () => (
  <div className="base">
    <header style={{display: 'flex', alignItems: 'center', justifyContent: 'center'}}>
      <p>Credit Report App</p>
    </header>
    <div className="container">
      <Route path="/" exact component={CreateUser}/>
      <Route path="/questions" component={Questions}/>
      <Route path="/order" component={OrderReport}/>
    </div>
  </div>
);

function CreateUser(props) {
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [ssn, setSSN] = useState('');
  const [dob, setDOB] = useState('');
  const [street, setStreet] = useState('');
  const [city, setCity] = useState('');
  const [state, setState] = useState('');
  const [zip, setZip] = useState(0);
  const history = useHistory();
  const { setUserDetails } = useContext(UserContext);

  function handleSubmit() {

    const appKey = uuidv4();
    let data = {
      appKey: appKey,
      firstName: firstName,
      lastName: lastName,
      ssn: ssn,
      dob: dob,
      address: {
        state: state,
        street: street,
        city: city,
        zip: zip,
      }
    };
    instance
      .post("/register", data)
      .then(function (response) {
        console.log("response,", response.data);
        setUserDetails(data);
        history.push({
          pathname: '/questions',
          state: {
            appKey: appKey,
            clientId: response.data.clientKey,
            userId: response.data.userId,
            authToken: response.data.authToken
          }
        });
      })
      .catch(function (error) {
        setUserDetails({});
        alert("User Creation Failed");
      });
  }

  return (
    <div style={{display: 'flex', alignItems: 'center', justifyContent: 'center'}}>
      <table>
        <tr>
          <th>Create Account</th>
        </tr>
        <tr>
          <td>First name *</td>
          <td>
            <input
              type="text"
              id="firstName"
              name="firstName"
              value={firstName}
              required
              onChange={(e) => setFirstName(e.target.value)}
            />
          </td>
        </tr>
        <tr>
          <td>Last name *</td>
          <td>
            <input
              type="text"
              id="lastName"
              name="lastName"
              value={lastName}
              required
              onChange={(e) => setLastName(e.target.value)}
            />
          </td>
        </tr>
        <tr>
          <td>SSN</td>
          <td>
            <input
              type="SSN"
              id="lastName"
              name="lastName"
              value={ssn}
              onChange={(e) => setSSN(e.target.value)}
            />
          </td>
        </tr>
        <tr>
          <td>DOB</td>
          <td>
            <input
              type="date"
              id="dob"
              name="dob"
              value={dob}
              onChange={(e) => setDOB(e.target.value)}
            />
          </td>
        </tr>
        <tr>
          <td>Street</td>
          <td>
            <input
              type="text"
              required
              id="street"
              name="street"
              value={street}
              onChange={(e) => setStreet(e.target.value)}
            />
          </td>
        </tr>
        <tr>
          <td>City</td>
          <td>
            <input
              type="text"
              required
              id="city"
              name="city"
              value={city}
              onChange={(e) => setCity(e.target.value)}
            />
          </td>
        </tr>
        <tr>
          <td>State</td>
          <td>
            <input
              type="text"
              required
              id="state"
              name="state"
              value={state}
              onChange={(e) => setState(e.target.value)}
            />
          </td>
        </tr>
        <tr>
          <td>ZIP</td>
          <td>
            <input
              type="number"
              required
              id="zip"
              name="zip"
              value={zip}
              onChange={(e) => setZip(e.target.value)}
            />
          </td>
        </tr>
        <tr>
          <td></td>
          <td>
            <input
              type="button"
              id="createUser"
              value="Create account"
              onClick={handleSubmit}
            />
          </td>
        </tr>
        <tr>
          <td>
            <Link to="/signup">Sign-up</Link>
          </td>
        </tr>
      </table>
    </div>
  );
}

const Questions = () => {
  const [questions, setQuestions] = useState([]);
  const questionAndAnswers = {};
  const { state } = useLocation();
  const history = useHistory();
  const { userDetails } = useContext(UserContext);

  const getQuestions = () => {
    instance.get('/questions').then(r => {
      setQuestions(r.data);
    }).catch(err => {
      alert(err);
    })
  }

  useEffect(() => {
    getQuestions();
  }, []);

  const onChangeAnswer = (questionIndex, answerIndex) => {
    questionAndAnswers[`${questionIndex}`] = answerIndex;
    console.log(questionAndAnswers);
  }

  const SaveAnswers = () => {
    instance.post("/questions/submit", {
      clientKey: state.clientId,
      appKey: state.appKey,
      authToken: state.authToken.replace(/\n/g, ''),
      answers: questionAndAnswers
    }).then(r => {
      alert(r.data.Message);
      history.push({
        pathname: '/order/report',
        state: {
          appKey: state.appKey,
          clientId: state.clientId,
          userId: state.userId,
          authToken: state.authToken.replace(/\n/g, '')
        }
      });
    }).catch(err => {
      alert(err);
    })
  };

  const onSubmit = () => {
    if (Object.keys(questionAndAnswers).length === questions.length) {
      SaveAnswers();
    } else {
      alert('you have to choose all the questions');
    }
  }

  return (
    <div style={{display: 'flex', alignItems: 'center', justifyContent: 'center'}}>
      <div>
        <p><b>Hello {userDetails.firstName}</b></p>
        {
          questions.map((question, index) => {
            return (
              <div>
                <p>{question.text}</p>
                  <select onChange={(e) => {
                    if (e.target.value) {
                      onChangeAnswer(question.id, e.target.value)
                    }
                  }}>
                    <option value={''}>select</option>
                  {
                    question.answers.map(answer => {
                      return  (
                        <option value={answer.id}>{answer.text}</option>
                      )
                    })
                  }
                  </select>
              </div>
            )
          })
        }
        <input type='button' value='Submit' onClick={onSubmit} />
      </div>
    </div>
  )
};

const OrderReport = () => {
  const products = [
    "AIRPLANE",
    "ATTORNEY",
    "AUTO LEASE",
    "Auto Loan",
    "Refinance",
    "BOAT",
    "BUSINESS CC",
    "Business",
  ];
  const { state } = useLocation();
  console.log('this is code', state);
  const [productCode, setProductCode] = useState('');
  const [displayToken, setDisplayToken] = useState('');
  const [reportKey, setReportKey] = useState('');
  const { userDetails } = useContext(UserContext);

  const OrderProduct = () => {
    instance.post('/order', {
      productCode: productCode,
      clientKey: state.clientId
    }).then(r => {
      setDisplayToken(r.data.displayToken);
      setReportKey(r.data.reportKey);
      alert('Thank you for ordering credit report \n Your report will be generated in sometime');
    }).catch(err => {
      alert(err);
    })
  }

  return (
    <div style={{display: 'flex', alignItems: 'center', justifyContent: 'center'}}>
      <div>
        <p><b>Hello {userDetails.firstName}</b></p>
        <label>Product Code</label>
        <select onChange={e => setProductCode(e.target.value)}>
          {
            products.map((item) => {
              return (<option value={item}>{item}</option>)
            })
          }
        </select>
        <input type='button' value='submit' onClick={() => OrderProduct()}/>

        {
          displayToken ? (
            <div>
              <label>Display token: </label>
              <p>{displayToken}</p>
            </div>
          ) : null
        }
        {
          reportKey ? (
            <div>
              <label>Report token: </label>
              <p>{reportKey}</p>
            </div>
          ) : null
        }

      </div>
    </div>
  )
};

export default App;
