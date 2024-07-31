import React, { useCallback } from 'react';
import {
    Alert,
    Bullseye,
    Button,
    Checkbox,
    Divider,
    DropEvent,
    Flex,
    FlexItem,
    Popover,
    Spinner,
    Stack,
    StackItem,
    Tab,
    TabContent,
    Tabs,
    TabTitleText,
    Text,
    TextContent,
    TextVariants,
    Title,
} from '@patternfly/react-core';
import { HelpIcon } from '@patternfly/react-icons';
import sortBy from 'lodash/sortBy';

import useRestQuery from 'hooks/useRestQuery';
import useAnalytics, { GENERATE_NETWORK_POLICIES } from 'hooks/useAnalytics';
import useURLSearch from 'hooks/useURLSearch';
import useTabs from 'hooks/patternfly/useTabs';
import { getRequestQueryStringForSearchFilter } from 'utils/searchUtils';
import { getQueryObject, getQueryString } from 'utils/queryStringUtils';
import { fetchNetworkPoliciesByClusterId } from 'services/NetworkService';

import ViewActiveYAMLs from './ViewActiveYAMLs';
import {
    NetworkPolicySimulator,
    SetNetworkPolicyModification,
} from '../hooks/useNetworkPolicySimulator';
import NetworkPoliciesYAML from './NetworkPoliciesYAML';
import {
    getDisplayYAMLFromNetworkPolicyModification,
    getSearchFilterFromScopeHierarchy,
} from '../utils/simulatorUtils';
import UploadYAMLButton from './UploadYAMLButton';
import NetworkSimulatorActions from './NetworkSimulatorActions';
import NotifyYAMLModal from './NotifyYAMLModal';
import { NetworkScopeHierarchy } from '../types/networkScopeHierarchy';
import CompareYAMLModal from './CompareYAMLModal';
import CodeCompareIcon from './CodeCompareIcon';
import NetworkPoliciesGenerationScope from './NetworkPoliciesGenerationScope';
import { getPropertiesForAnalytics } from '../utils/networkGraphURLUtils';

// @TODO: Consider a better approach to managing the side panel related state (simulation + URL path for entities)
export function clearSimulationQuery(search: string): string {
    const modifiedSearchFilter = getQueryObject(search);
    delete modifiedSearchFilter.simulation;
    const queryString = getQueryString(modifiedSearchFilter);
    return queryString;
}

export type NetworkPolicySimulatorSidePanelProps = {
    simulator: NetworkPolicySimulator;
    setNetworkPolicyModification: SetNetworkPolicyModification;
    /** scopeHierarchy is the user's selected scope for the network graph */
    scopeHierarchy: NetworkScopeHierarchy;
    scopeDeploymentCount: number;
};

const tabs = {
    SIMULATE_NETWORK_POLICIES: 'Simulate network policies',
    VIEW_ACTIVE_YAMLS: 'View active YAMLS',
}; // space in visible text, but id has underscore!

function NetworkPolicySimulatorSidePanel({
    simulator,
    setNetworkPolicyModification,
    scopeHierarchy,
    scopeDeploymentCount,
}: NetworkPolicySimulatorSidePanelProps) {
    const { analyticsTrack } = useAnalytics();
    const { searchFilter } = useURLSearch();

    const { activeKeyTab, onSelectTab } = useTabs({
        defaultTab: tabs.SIMULATE_NETWORK_POLICIES,
    });
    const [isExcludingPortsAndProtocols, setIsExcludingPortsAndProtocols] =
        React.useState<boolean>(false);
    const [isNotifyModalOpen, setIsNotifyModalOpen] = React.useState(false);
    const [compareModalYAMLs, setCompareModalYAMLs] = React.useState<{
        generated: string;
        current: string;
    } | null>(null);

    const clusterId = scopeHierarchy.cluster.id;
    const deploymentQuery = getRequestQueryStringForSearchFilter(
        getSearchFilterFromScopeHierarchy(scopeHierarchy)
    );

    const fetchNetworkPolicies = useCallback(
        () =>
            fetchNetworkPoliciesByClusterId(clusterId, deploymentQuery).then((policies) =>
                sortBy(policies, 'name')
            ),
        [clusterId, deploymentQuery]
    );
    const { data: currentNetworkPolicies } = useRestQuery(fetchNetworkPolicies);

    function handleFileInputChange(_event: DropEvent, file: File) {
        if (file && !file.name.includes('.yaml')) {
            setNetworkPolicyModification({
                state: 'UPLOAD',
                options: {
                    modification: null,
                    error: 'File must be .yaml',
                },
            });
        } else {
            const reader = new FileReader();
            reader.onload = () => {
                const fileAsBinaryString = reader.result;
                setNetworkPolicyModification({
                    state: 'UPLOAD',
                    options: {
                        modification: {
                            applyYaml: fileAsBinaryString as string,
                            toDelete: [],
                        },
                        error: '',
                    },
                });
            };
            reader.onerror = () => {
                setNetworkPolicyModification({
                    state: 'UPLOAD',
                    options: {
                        modification: null,
                        error: 'Could not read file',
                    },
                });
            };
            reader.readAsBinaryString(file);
        }
    }

    function generateNetworkPolicies() {
        const properties = getPropertiesForAnalytics(searchFilter);

        analyticsTrack({
            event: GENERATE_NETWORK_POLICIES,
            properties,
        });
        setNetworkPolicyModification({
            state: 'GENERATED',
            options: {
                scopeHierarchy,
                networkDataSince: '',
                excludePortsAndProtocols: isExcludingPortsAndProtocols,
            },
        });
    }

    function undoNetworkPolicies() {
        setNetworkPolicyModification({
            state: 'UNDO',
            options: {
                clusterId: scopeHierarchy.cluster.id,
            },
        });
    }

    function openNotifyYAMLModal() {
        setIsNotifyModalOpen(true);
    }

    if (simulator.isLoading) {
        return (
            <Bullseye>
                <Spinner size="lg" />
            </Bullseye>
        );
    }

    if (simulator.state === 'GENERATED') {
        const currentPolicies = currentNetworkPolicies ?? [];
        const currentYaml =
            currentPolicies.length === 0
                ? 'No network policies exist in the current scope'
                : currentPolicies.map((policy) => policy.yaml).join('\n---\n');
        const generatedYaml = getDisplayYAMLFromNetworkPolicyModification(simulator.modification);
        return (
            <div>
                <Flex
                    direction={{ default: 'column' }}
                    spaceItems={{ default: 'spaceItemsSm' }}
                    className="pf-v5-u-p-md pf-v5-u-pb-sm pf-v5-u-mb-0"
                >
                    <FlexItem>
                        <TextContent>
                            <Text component={TextVariants.h2}>Generated network policies</Text>
                            <Text component={TextVariants.h3} className="pf-v5-u-m-0">
                                Scope of baseline:
                            </Text>
                        </TextContent>
                    </FlexItem>
                    <NetworkPoliciesGenerationScope
                        scopeHierarchy={scopeHierarchy}
                        scopeDeploymentCount={scopeDeploymentCount}
                    />
                </Flex>
                <Flex direction={{ default: 'column' }}>
                    <FlexItem style={{ overflow: 'auto' }}>
                        <NetworkPoliciesYAML
                            yaml={generatedYaml}
                            additionalControls={[
                                <Flex
                                    justifyContent={{ default: 'justifyContentFlexEnd' }}
                                    alignItems={{ default: 'alignItemsCenter' }}
                                    spaceItems={{ default: 'spaceItemsNone' }}
                                    className="pf-v5-u-flex-1"
                                >
                                    <Button
                                        variant="link"
                                        onClick={() =>
                                            setCompareModalYAMLs({
                                                generated: generatedYaml,
                                                current: currentYaml,
                                            })
                                        }
                                        icon={<CodeCompareIcon />}
                                    >
                                        Compare
                                    </Button>
                                    <Popover
                                        bodyContent={
                                            <Flex spaceItems={{ default: 'spaceItemsSm' }}>
                                                <Title headingLevel="h3">Compare</Title>
                                                <Text>
                                                    Compare the generated network policies to the
                                                    existing network policies.
                                                </Text>
                                            </Flex>
                                        }
                                    >
                                        <button
                                            className="pf-v5-u-color-200"
                                            type="button"
                                            aria-label="More info on comparing changes"
                                        >
                                            <HelpIcon />
                                        </button>
                                    </Popover>
                                </Flex>,
                            ]}
                        />
                    </FlexItem>
                    <FlexItem>
                        <NetworkSimulatorActions
                            generateNetworkPolicies={generateNetworkPolicies}
                            undoNetworkPolicies={undoNetworkPolicies}
                            onFileInputChange={handleFileInputChange}
                            openNotifyYAMLModal={openNotifyYAMLModal}
                        />
                    </FlexItem>
                </Flex>
                <NotifyYAMLModal
                    isModalOpen={isNotifyModalOpen}
                    setIsModalOpen={setIsNotifyModalOpen}
                    clusterId={scopeHierarchy.cluster.id}
                    modification={simulator.modification}
                />
                {compareModalYAMLs && (
                    <CompareYAMLModal
                        generated={compareModalYAMLs.generated}
                        current={compareModalYAMLs.current}
                        isOpen={!!compareModalYAMLs}
                        onClose={() => setCompareModalYAMLs(null)}
                    />
                )}
            </div>
        );
    }

    // @TODO: Consider how to reuse parts of this that are similar between states
    if (simulator.state === 'UNDO') {
        const yaml = getDisplayYAMLFromNetworkPolicyModification(simulator.modification);
        return (
            <div>
                <Flex
                    direction={{ default: 'row' }}
                    alignItems={{ default: 'alignItemsFlexEnd' }}
                    className="pf-v5-u-p-lg pf-v5-u-mb-0"
                >
                    <FlexItem>
                        <TextContent>
                            <Text component={TextVariants.h2} className="pf-v5-u-font-size-xl">
                                Network Policy Simulator
                            </Text>
                        </TextContent>
                    </FlexItem>
                </Flex>
                <Divider component="div" />
                <Stack hasGutter>
                    <StackItem className="pf-v5-u-p-md">
                        <Alert
                            variant={simulator.error ? 'danger' : 'success'}
                            isInline
                            isPlain
                            title={
                                simulator.error
                                    ? simulator.error
                                    : 'Viewing modification that will undo last applied change'
                            }
                            component="p"
                        />
                    </StackItem>
                    <StackItem isFilled style={{ overflow: 'auto' }}>
                        <NetworkPoliciesYAML yaml={yaml} />
                    </StackItem>
                    <StackItem>
                        <NetworkSimulatorActions
                            generateNetworkPolicies={generateNetworkPolicies}
                            undoNetworkPolicies={undoNetworkPolicies}
                            onFileInputChange={handleFileInputChange}
                            openNotifyYAMLModal={openNotifyYAMLModal}
                        />
                    </StackItem>
                </Stack>
                <NotifyYAMLModal
                    isModalOpen={isNotifyModalOpen}
                    setIsModalOpen={setIsNotifyModalOpen}
                    clusterId={scopeHierarchy.cluster.id}
                    modification={simulator.modification}
                />
            </div>
        );
    }

    if (simulator.state === 'UPLOAD') {
        const yaml = getDisplayYAMLFromNetworkPolicyModification(simulator.modification);
        return (
            <div>
                <Flex
                    direction={{ default: 'row' }}
                    alignItems={{ default: 'alignItemsFlexEnd' }}
                    className="pf-v5-u-p-lg pf-v5-u-mb-0"
                >
                    <FlexItem>
                        <TextContent>
                            <Text component={TextVariants.h2} className="pf-v5-u-font-size-xl">
                                Network Policy Simulator
                            </Text>
                        </TextContent>
                    </FlexItem>
                </Flex>
                <Divider component="div" />
                <Stack hasGutter>
                    <StackItem className="pf-v5-u-p-md">
                        <Alert
                            variant={simulator.error ? 'danger' : 'success'}
                            isInline
                            isPlain
                            title={
                                simulator.error ? simulator.error : 'Uploaded policies processed'
                            }
                            component="p"
                        />
                    </StackItem>
                    <StackItem isFilled style={{ overflow: 'auto' }}>
                        <NetworkPoliciesYAML yaml={yaml} />
                    </StackItem>
                    <StackItem>
                        <NetworkSimulatorActions
                            generateNetworkPolicies={generateNetworkPolicies}
                            undoNetworkPolicies={undoNetworkPolicies}
                            onFileInputChange={handleFileInputChange}
                            openNotifyYAMLModal={openNotifyYAMLModal}
                        />
                    </StackItem>
                </Stack>
                <NotifyYAMLModal
                    isModalOpen={isNotifyModalOpen}
                    setIsModalOpen={setIsNotifyModalOpen}
                    clusterId={scopeHierarchy.cluster.id}
                    modification={simulator.modification}
                />
            </div>
        );
    }

    return (
        <Stack>
            <StackItem>
                <Flex
                    direction={{ default: 'column' }}
                    spaceItems={{ default: 'spaceItemsSm' }}
                    className="pf-v5-u-p-md pf-v5-u-pb-sm pf-v5-u-mb-0"
                >
                    <TextContent>
                        <Text component={TextVariants.h2}>Generate network policies</Text>
                        <Text component={TextVariants.h3} className="pf-v5-u-m-0">
                            Scope of baseline:
                        </Text>
                    </TextContent>
                    <NetworkPoliciesGenerationScope
                        scopeHierarchy={scopeHierarchy}
                        scopeDeploymentCount={scopeDeploymentCount}
                    />
                </Flex>
            </StackItem>
            <StackItem>
                <Tabs activeKey={activeKeyTab} onSelect={onSelectTab}>
                    <Tab
                        eventKey={tabs.SIMULATE_NETWORK_POLICIES}
                        tabContentId="Simulate_network_policies"
                        title={<TabTitleText>{tabs.SIMULATE_NETWORK_POLICIES}</TabTitleText>}
                    />
                    <Tab
                        eventKey={tabs.VIEW_ACTIVE_YAMLS}
                        tabContentId="View_active_YAMLS"
                        title={<TabTitleText>{tabs.VIEW_ACTIVE_YAMLS}</TabTitleText>}
                    />
                </Tabs>
            </StackItem>
            <StackItem isFilled style={{ overflow: 'auto' }}>
                <TabContent
                    eventKey={tabs.SIMULATE_NETWORK_POLICIES}
                    id="Simulate_network_policies"
                    hidden={activeKeyTab !== tabs.SIMULATE_NETWORK_POLICIES}
                >
                    <div className="pf-v5-u-p-lg pf-v5-u-h-100">
                        <Stack hasGutter>
                            <StackItem>
                                <Stack hasGutter>
                                    <StackItem>
                                        <TextContent>
                                            <Text
                                                component={TextVariants.h2}
                                                className="pf-v5-u-font-size-lg"
                                            >
                                                Generate network policies from the traffic
                                            </Text>
                                        </TextContent>
                                    </StackItem>
                                    <StackItem>
                                        <TextContent>
                                            <Text component={TextVariants.p}>
                                                Generate a set of recommended network policies based
                                                on your cluster&apos;s traffic. Only deployments
                                                that are part of the current scope will be included
                                                in generated policies.
                                            </Text>
                                        </TextContent>
                                    </StackItem>
                                    <StackItem>
                                        <Checkbox
                                            label="Exclude ports & protocols"
                                            isChecked={isExcludingPortsAndProtocols}
                                            onChange={(_event, val) =>
                                                setIsExcludingPortsAndProtocols(val)
                                            }
                                            id="controlled-check-1"
                                            name="check1"
                                        />
                                    </StackItem>
                                    <StackItem>
                                        <Button
                                            variant="secondary"
                                            onClick={generateNetworkPolicies}
                                        >
                                            Generate and simulate network policies
                                        </Button>
                                    </StackItem>
                                </Stack>
                            </StackItem>
                            <StackItem>
                                <Divider component="div" />
                            </StackItem>
                            <StackItem>
                                <Stack hasGutter>
                                    <StackItem>
                                        <TextContent>
                                            <Text
                                                component={TextVariants.h2}
                                                className="pf-v5-u-font-size-lg"
                                            >
                                                Upload a network policy YAML
                                            </Text>
                                        </TextContent>
                                    </StackItem>
                                    <StackItem>
                                        <TextContent>
                                            <Text component={TextVariants.p}>
                                                Upload your network policies to quickly preview your
                                                environment under different policy configurations
                                                and time windows. When ready, apply the network
                                                policies directly or share them with your team.
                                            </Text>
                                        </TextContent>
                                    </StackItem>
                                    <StackItem>
                                        <UploadYAMLButton
                                            onFileInputChange={handleFileInputChange}
                                        />
                                    </StackItem>
                                </Stack>
                            </StackItem>
                        </Stack>
                    </div>
                </TabContent>
                <TabContent
                    eventKey={tabs.VIEW_ACTIVE_YAMLS}
                    id="View_active_YAMLS"
                    hidden={activeKeyTab !== tabs.VIEW_ACTIVE_YAMLS}
                >
                    <ViewActiveYAMLs
                        networkPolicies={currentNetworkPolicies ?? []}
                        generateNetworkPolicies={generateNetworkPolicies}
                        undoNetworkPolicies={undoNetworkPolicies}
                        onFileInputChange={handleFileInputChange}
                    />
                </TabContent>
            </StackItem>
        </Stack>
    );
}

export default NetworkPolicySimulatorSidePanel;
