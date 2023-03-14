import { Modal, Table } from 'react-bootstrap';

function EventRulesModal(props) {
  const { showModal, onHide, trackingPlanDetail } = props;

  return (
    <Modal show={showModal} onHide={onHide}>
      <Modal.Header closeButton>
        <Modal.Title>Event Rules</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            {trackingPlanDetail.map((rule) => (
              <tr key={rule.id}>
                <td>{rule.id}</td>
                <td>{rule.name}</td>
                <td>{rule.description}</td>
              </tr>
            ))}
          </tbody>
        </Table>
      </Modal.Body>
    </Modal>
  );
}

export default EventRulesModal;
