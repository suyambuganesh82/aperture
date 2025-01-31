package com.fluxninja.aperture.armeria;

import com.fluxninja.aperture.sdk.*;
import com.linecorp.armeria.common.HttpRequest;
import com.linecorp.armeria.common.HttpResponse;
import com.linecorp.armeria.server.HttpService;
import com.linecorp.armeria.server.ServiceRequestContext;
import com.linecorp.armeria.server.SimpleDecoratingHttpService;
import java.util.Collections;
import java.util.Map;
import java.util.function.Function;

/** Decorates an {@link HttpService} to enable flow control using provided {@link ApertureSDK} */
public class ApertureHTTPService extends SimpleDecoratingHttpService {
    private final ApertureSDK apertureSDK;
    private final String controlPointName;
    private final boolean failOpen;

    public static Function<? super HttpService, ApertureHTTPService> newDecorator(
            ApertureSDK apertureSDK, String controlPointName) {
        ApertureHTTPServiceBuilder builder = new ApertureHTTPServiceBuilder();
        builder.setApertureSDK(apertureSDK).setControlPointName(controlPointName);
        return builder::build;
    }

    public static Function<? super HttpService, ApertureHTTPService> newDecorator(
            ApertureSDK apertureSDK, String controlPointName, boolean failOpen) {
        ApertureHTTPServiceBuilder builder = new ApertureHTTPServiceBuilder();
        builder.setApertureSDK(apertureSDK)
                .setControlPointName(controlPointName)
                .setEnableFailOpen(failOpen);
        return builder::build;
    }

    public ApertureHTTPService(
            HttpService delegate,
            ApertureSDK apertureSDK,
            String controlPointName,
            boolean failOpen) {
        super(delegate);
        this.apertureSDK = apertureSDK;
        this.controlPointName = controlPointName;
        this.failOpen = failOpen;
    }

    @Override
    public HttpResponse serve(ServiceRequestContext ctx, HttpRequest req) throws Exception {
        TrafficFlowRequest request =
                HttpUtils.trafficFlowRequestFromRequest(ctx, req, controlPointName);
        TrafficFlow flow = this.apertureSDK.startTrafficFlow(request);

        if (flow.ignored()) {
            return unwrap().serve(ctx, req);
        }

        FlowDecision flowDecision = flow.getDecision();
        boolean flowAccepted =
                (flowDecision == FlowDecision.Accepted
                        || (flowDecision == FlowDecision.Unreachable && this.failOpen));

        if (flowAccepted) {
            HttpResponse res;
            try {
                Map<String, String> newHeaders = Collections.emptyMap();
                if (flow.checkResponse() != null) {
                    newHeaders = flow.checkResponse().getOkResponse().getHeadersMap();
                }
                HttpRequest newRequest = HttpUtils.updateHeaders(req, newHeaders);
                ctx.updateRequest(newRequest);

                res = unwrap().serve(ctx, newRequest);
            } catch (Exception e) {
                flow.setStatus(FlowStatus.Error);
                throw e;
            } finally {
                flow.end();
            }
            return res;
        } else {
            return HttpUtils.handleRejectedFlow(flow);
        }
    }
}
