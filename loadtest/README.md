# Load test com k6 (in-cluster)

Teste de stress leve rodando **dentro do GKE**, batendo direto no `Service`
interno (`ClusterIP`). Sem egress, sem passar pelo Load Balancer — o mais barato
e o melhor pra observar o HPA (1 → 5 réplicas).

## Rodar

```bash
# 1. Cria/atualiza o ConfigMap a partir do script (fonte da verdade = script.js)
kubectl create configmap k6-script \
  --from-file=script.js=loadtest/script.js \
  --dry-run=client -o yaml | kubectl apply -f -

# 2. Aplica o Job (garante que nao existe um anterior)
kubectl delete job k6-loadtest --ignore-not-found
kubectl apply -f loadtest/k6-job.yaml

# 3. Acompanha o resultado do k6 em tempo real
kubectl logs -f job/k6-loadtest
```

## Observar o autoscaling em outro terminal

```bash
# HPA reagindo (CPU% e replicas subindo)
kubectl get hpa fraud-score -w

# Pods sendo criados/removidos
kubectl get pods -l app=fraud-score -w
```

## Custo — como não gastar

- O teste dura ~3min30 (ramp-up 1m + carga 2m + ramp-down 30s). Rajada curta.
- `ttlSecondsAfterFinished: 600` faz o Job se autolimpar.
- Depois de validar, o que gera custo contínuo é **o cluster ligado e o Load
  Balancer do Ingress**. Se for ambiente só de teste:

  ```bash
  kubectl delete -f k8s/manifests/ingress.yaml   # derruba o LB externo
  # ou, pra zerar tudo:
  cd infra/terraform && terraform destroy
  ```

## Ajustar a carga

Edite os `stages` em `loadtest/script.js` (ex.: `target: 50` → mais VUs) e
**regenere o ConfigMap** (passo 1). O ConfigMap dentro de `k6-job.yaml` é só um
placeholder — a fonte da verdade é sempre o `script.js`.
```
