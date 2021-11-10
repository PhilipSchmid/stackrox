import React, { ReactElement, useContext, useState } from 'react';
import { Card, InputGroup, Modal, ModalVariant, TextInput } from '@patternfly/react-core';
import { TableComposable, Tbody, Td, Th, Thead, Tr } from '@patternfly/react-table';

import workflowStateContext from 'Containers/workflowStateContext';
import { ComponentWhereCVEOccurs } from '../types';

export type ComponentsModalProps = {
    components: ComponentWhereCVEOccurs[];
    onClose: () => void;
};

function ComponentsModal({ components, onClose }: ComponentsModalProps): ReactElement {
    const workflowState = useContext(workflowStateContext);
    const [inputValue, setInputValue] = useState('');

    function onInputValueChange(value) {
        setInputValue(value);
    }

    const filteredComponents = components.filter((component) => {
        return component.name.includes(inputValue);
    });

    return (
        <Modal
            variant={ModalVariant.small}
            title="Affected Components"
            isOpen={components.length !== 0}
            onClose={onClose}
        >
            <InputGroup className="pf-u-mt-md">
                <TextInput
                    name="componentsFilter"
                    id="componentsFilter"
                    type="text"
                    aria-label="Filter components"
                    placeholder="Filter components"
                    value={inputValue}
                    onChange={onInputValueChange}
                />
            </InputGroup>
            <Card isFlat className="pf-u-mt-lg">
                <TableComposable aria-label="Components Table" variant="compact" borders>
                    <Thead>
                        <Tr>
                            <Th>Component</Th>
                            <Th>Fixed in</Th>
                        </Tr>
                    </Thead>
                    <Tbody>
                        {filteredComponents.map((component) => {
                            const componentURL = workflowState
                                .pushList('COMPONENT')
                                .pushListItem(component.id)
                                .toUrl();

                            return (
                                <Tr key={component.name}>
                                    <Td dataLabel="Component">
                                        <a href={componentURL} target="_blank" rel="noreferrer">
                                            {component.name}
                                        </a>
                                    </Td>
                                    <Td dataLabel="Fixed in">{component.fixedIn}</Td>
                                </Tr>
                            );
                        })}
                    </Tbody>
                </TableComposable>
            </Card>
        </Modal>
    );
}

export default ComponentsModal;
