package com.github.kr328.clash.service;

import com.github.kr328.clash.service.IClashEventObserver;
import com.github.kr328.clash.service.IClashEventService;
import com.github.kr328.clash.service.IClashProfileService;
import com.github.kr328.clash.service.IClashSettingService;
import com.github.kr328.clash.callback.IUrlTestCallback;
import com.github.kr328.clash.core.event.Event;
import com.github.kr328.clash.core.model.Packet;

interface IClashService {
    // Services
    IClashEventService getEventService();
    IClashProfileService getProfileService();
    IClashSettingService getSettingService();

    // Status
    ProcessEvent getCurrentProcessStatus();

    // Control
    void setSelectProxy(String proxy, String selected);
    void startTunDevice(in ParcelFileDescriptor fd, int mtu, boolean dnsHijacking);
    void stopTunDevice();
    void start();
    void stop();

    // Query
    ProxyPacket queryAllProxies();
    GeneralPacket queryGeneral();
    void startUrlTest(in String[] proxies, IUrlTestCallback callback);
}
