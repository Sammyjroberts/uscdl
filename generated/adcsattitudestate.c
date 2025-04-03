/**
* ADCSAttitudeState
* Current attitude state of the spacecraft
*/

#include "adcsattitudestate.h"
#include <string.h>
#include <stdlib.h>

void adcs_attitude_state_init(ADCSAttitudeState_t* p_data) {
    if (p_data == NULL) {
        return;
    }
    memset(p_data->quaternion, 0.0, sizeof(p_data->quaternion));
    memset(p_data->angularVelocity, 0.0, sizeof(p_data->angularVelocity));
    p_data->timestamp = 0;
    p_data->attitudeDeterminationMode = 0;
    p_data->attitudeValid = false;
}

int adcs_attitude_state_serialize(const ADCSAttitudeState_t* p_data, uint8_t* buffer, size_t buffer_size) {
    if (p_data == NULL || buffer == NULL) {
        return -1;
    }

    // Ensure buffer is large enough
    if (buffer_size < 34) {
        return -1;
    }

    size_t offset = 0;
    uint8_t* ptr = buffer;
    size_t item_size = 0;
    // Direct copy for little-endian or byte types
    item_size = 4 * 4;
    memcpy(ptr + offset, p_data->quaternion, item_size);
    offset += item_size;
    // Direct copy for little-endian or byte types
    item_size = 4 * 3;
    memcpy(ptr + offset, p_data->angularVelocity, item_size);
    offset += item_size;
    // Direct copy for little-endian or byte types
    memcpy(ptr + offset, &p_data->timestamp, 4);
    offset += 4;
    // Direct copy for little-endian or byte types
    memcpy(ptr + offset, &p_data->attitudeDeterminationMode, 1);
    offset += 1;
    // Direct copy for little-endian or byte types
    memcpy(ptr + offset, &p_data->attitudeValid, 1);
    offset += 1;

    return (int)offset;
}

int adcs_attitude_state_deserialize(ADCSAttitudeState_t* p_data, const uint8_t* buffer, size_t buffer_size) {
    if (p_data == NULL || buffer == NULL) {
        return -1;
    }

    // Initialize the structure
    adcs_attitude_state_init(p_data);

    size_t offset = 0;
    const uint8_t* ptr = buffer;
    size_t item_size = 0;
    // Direct copy for little-endian or byte types
    item_size = 4 * 4;
    if (offset + item_size > buffer_size) {
        return -1;
    }
    memcpy(p_data->quaternion, ptr + offset, item_size);
    offset += item_size;
    // Direct copy for little-endian or byte types
    item_size = 4 * 3;
    if (offset + item_size > buffer_size) {
        return -1;
    }
    memcpy(p_data->angularVelocity, ptr + offset, item_size);
    offset += item_size;
    // Direct copy for little-endian or byte types
    if (offset + 4 > buffer_size) {
        return -1;
    }
    memcpy(&p_data->timestamp, ptr + offset, 4);
    offset += 4;
    // Direct copy for little-endian or byte types
    if (offset + 1 > buffer_size) {
        return -1;
    }
    memcpy(&p_data->attitudeDeterminationMode, ptr + offset, 1);
    offset += 1;
    // Direct copy for little-endian or byte types
    if (offset + 1 > buffer_size) {
        return -1;
    }
    memcpy(&p_data->attitudeValid, ptr + offset, 1);
    offset += 1;

    return (int)offset;
}
