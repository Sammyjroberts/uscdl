/**
* ADCSActuatorCommands
* Commands for the attitude control actuators
*/

#include "adcsactuatorcommands.h"
#include <string.h>
#include <stdlib.h>

void adcs_actuator_commands_init(ADCSActuatorCommands_t* p_data) {
    if (p_data == NULL) {
        return;
    }
    memset(p_data->reactionWheelSpeeds, 0, sizeof(p_data->reactionWheelSpeeds));
    memset(p_data->magnetorquerCommands, 0, sizeof(p_data->magnetorquerCommands));
    p_data->commandTimestamp = 0;
    p_data->controlMode = 0;
}

int adcs_actuator_commands_serialize(const ADCSActuatorCommands_t* p_data, uint8_t* buffer, size_t buffer_size) {
    if (p_data == NULL || buffer == NULL) {
        return -1;
    }

    // Ensure buffer is large enough
    if (buffer_size < 19) {
        return -1;
    }

    size_t offset = 0;
    uint8_t* ptr = buffer;
    size_t item_size = 0;
    // Direct copy for little-endian or byte types
    item_size = 2 * 4;
    memcpy(ptr + offset, p_data->reactionWheelSpeeds, item_size);
    offset += item_size;
    // Direct copy for little-endian or byte types
    item_size = 2 * 3;
    memcpy(ptr + offset, p_data->magnetorquerCommands, item_size);
    offset += item_size;
    // Direct copy for little-endian or byte types
    memcpy(ptr + offset, &p_data->commandTimestamp, 4);
    offset += 4;
    // Direct copy for little-endian or byte types
    memcpy(ptr + offset, &p_data->controlMode, 1);
    offset += 1;

    return (int)offset;
}

int adcs_actuator_commands_deserialize(ADCSActuatorCommands_t* p_data, const uint8_t* buffer, size_t buffer_size) {
    if (p_data == NULL || buffer == NULL) {
        return -1;
    }

    // Initialize the structure
    adcs_actuator_commands_init(p_data);

    size_t offset = 0;
    const uint8_t* ptr = buffer;
    size_t item_size = 0;
    // Direct copy for little-endian or byte types
    item_size = 2 * 4;
    if (offset + item_size > buffer_size) {
        return -1;
    }
    memcpy(p_data->reactionWheelSpeeds, ptr + offset, item_size);
    offset += item_size;
    // Direct copy for little-endian or byte types
    item_size = 2 * 3;
    if (offset + item_size > buffer_size) {
        return -1;
    }
    memcpy(p_data->magnetorquerCommands, ptr + offset, item_size);
    offset += item_size;
    // Direct copy for little-endian or byte types
    if (offset + 4 > buffer_size) {
        return -1;
    }
    memcpy(&p_data->commandTimestamp, ptr + offset, 4);
    offset += 4;
    // Direct copy for little-endian or byte types
    if (offset + 1 > buffer_size) {
        return -1;
    }
    memcpy(&p_data->controlMode, ptr + offset, 1);
    offset += 1;

    return (int)offset;
}
