import { Container } from 'react-bootstrap'

const Header = ({ head, description }) => {
  return (
    <Container>
      <div className='starter-template text-center mt-5'>
        <h1>Rudderstack</h1>
        <hr></hr>
        <h3>{head} Page</h3>
        <p className='text-capitalize'>{description}</p>
      </div>
    </Container>
  )
}

export default Header